package service

type StorageInterface interface {
	ProjectStorageInterface
	UserStorageInterface
	HeroStorageInterface
	ProjSectDescStorageInterface
}

type Service struct {
	Project
	User
	Partner
	Hero
	ProjSectDesc
}

func NewService(storage StorageInterface) (Service, error) {
	return Service{
		Project{Storage: storage},
		User{Storage: storage},
		Partner{Storage: storage},
		Hero{Storage: storage},
		ProjSectDesc{Storage: storage},
	}, nil
}
