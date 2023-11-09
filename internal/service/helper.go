package service

import (
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strconv"
)

const (
	patternForFeedback = "From: Info@walking-school.site\r\n" +
		"Subject: Звернення на Школа-Ходи \r\n" +
		"\r\n" +
		"Ім'я: %s\r\n" +
		"Прізвище: %s\r\n" +
		"Телефон: %s\r\n" +
		"Пошта: %s\r\n" +
		"Повідомлення: %s\r\n"
	linkToResetPassword  = `https://walking-school.site/reset?token=%s`
	resetPasswordMessage = "From: Info@walking-school.site\r\n" +
		"Subject: Відновлення пароля облікового запису Школа-Ходи \r\n" +
		"\r\n" +
		"Ви або хтось ще запросили відновлення пароля для вашого облікового запису Школа-Ходи. Для завершення процесу відновлення пароля, будь ласка, дотримуйтесь інструкцій нижче.\r\n" +
		"1. Перейдіть за наступним посиланням для введення нового пароля: %s\r\n" +
		"2. Після переходу на посилання введіть новий пароль двічі, щоб підтвердити його.\r\n" +
		"Якщо ви не попросили відновлення пароля, проігноруйте це повідомлення. Ваш поточний пароль залишиться без змін.\r\n" +
		"З повагою, команда підтримки Школа-Ходи\r\n"
)

func SHA256(password, salt string) string {
	sum := sha256.Sum256([]byte(password + salt))

	return fmt.Sprintf("%x", sum)
}

func sendMessage(host, portS, username, password, from, to, msg string) error {
	portD, err := strconv.Atoi(portS)
	if err != nil {
		return fmt.Errorf("error occurred in strconv.Atoi(): %w", err)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", host, portD), tlsConfig)
	if err != nil {
		return fmt.Errorf("error occured in Dial(): %w", err)
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("error occured in NewClient(): %w", err)
	}

	auth := smtp.PlainAuth("", username, password, host)

	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("error occured in Auth(): %w", err)
	}

	if err = client.Mail(from); err != nil {
		return fmt.Errorf("error occured in Mail(): %w", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("error occured in Rcpt(): %w", err)
	}

	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("error occured in Data(): %w", err)
	}
	defer w.Close()

	_, err = w.Write([]byte(msg))
	if err != nil {
		return fmt.Errorf("error occured in Write(): %w", err)
	}

	return nil
}
