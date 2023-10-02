package service

type StorageInterface interface {
	ProjectInterface
}

type Service struct {
	Project
	User
	Partner
}

func NewService(storage StorageInterface) (Service, error) {
	return Service{
		Project{Storage: storage},
		User{Storage: storage},
		Partner{Storage: storage},
	}, nil
}
