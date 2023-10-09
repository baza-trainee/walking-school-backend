package service

type StorageInterface interface {
	ProjectStorageInterface
	UserStorageInterface
	HeroStorageInterface
	ProjSectDescStorageInterface
	PartnerStorageInterface
	ImagesCarouselStorageInterface
	ContactStorageInterface
}

type Service struct {
	Project
	User
	Partner
	Hero
	ProjSectDesc
	ImagesCarousel
	Contact
}

func NewService(storage StorageInterface) (Service, error) {
	return Service{
		Project{Storage: storage},
		User{Storage: storage},
		Partner{Storage: storage},
		Hero{Storage: storage},
		ProjSectDesc{Storage: storage},
		ImagesCarousel{Storage: storage},
		Contact{Storage: storage},
	}, nil
}
