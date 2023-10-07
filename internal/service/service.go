package service

type StorageInterface interface {
	ProjectStorageInterface
	UserStorageInterface
	HeroStorageInterface
}

type Service struct {
	Project
	User
	Partner
	Hero
}

func NewService(storage StorageInterface) (Service, error) {
	return Service{
		Project{Storage: storage},
		User{Storage: storage},
		Partner{Storage: storage},
		Hero{Storage: storage},
	}, nil
}
