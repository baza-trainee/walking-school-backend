package service

import (
	"github.com/baza-trainee/walking-school-backend/internal/config"
)

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
	Feedback
}

func NewService(storage StorageInterface, form config.Form) (Service, error) {
	return Service{
		Project{Storage: storage},
		User{Storage: storage},
		Partner{Storage: storage},
		Hero{Storage: storage},
		ProjSectDesc{Storage: storage},
		ImagesCarousel{Storage: storage},
		Contact{Storage: storage},
		Feedback{Cfg: form},
	}, nil
}
