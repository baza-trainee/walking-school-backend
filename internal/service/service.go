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
	AuthorizationStorageInterface
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
	Authorization
}

func NewService(storage StorageInterface, cfg config.Config) (Service, error) {
	return Service{
		Project{Storage: storage},
		User{Storage: storage},
		Partner{Storage: storage},
		Hero{Storage: storage},
		ProjSectDesc{Storage: storage},
		ImagesCarousel{Storage: storage},
		Contact{Storage: storage},
		Feedback{Cfg: cfg.Form},
		Authorization{Storage: storage, Cfg: cfg.Auth},
	}, nil
}
