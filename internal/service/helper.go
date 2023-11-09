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
		"Name: %s\r\n" +
		"Surname: %s\r\n" +
		"Phone: %s\r\n" +
		"Email: %s\r\n" +
		"Text: %s\r\n"
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
