package service

type ProjectInterface interface {
}

type Project struct {
	Storage ProjectInterface
}
