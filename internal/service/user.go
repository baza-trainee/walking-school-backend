package service

type UserInterface interface {
}

type User struct {
	Storage UserInterface
}
