package service

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/google/uuid"
)

type ImagesCarouselStorageInterface interface {
	CreateImagesCarouselStorage(context.Context, model.ImageCarousel) error
	GetAllImagesCarouselStorage(context.Context) ([]model.ImageCarousel, error)
	UpdateImagesCarouselByIDStorage(context.Context, model.ImageCarousel) error
	DeleteImagesCarouselByIDStorage(context.Context, string) error
}

type ImagesCarousel struct {
	Storage ImagesCarouselStorageInterface
}

func (ic ImagesCarousel) CreateImagesCarouselService(ctx context.Context, imagesCarousel model.ImageCarousel) error {
	imagesCarousel.ID = uuid.NewString()

	if err := ic.Storage.CreateImagesCarouselStorage(ctx, imagesCarousel); err != nil {
		return fmt.Errorf("error occurred in CreateImagesCarouselStorage: %w", err)
	}

	return nil
}

func (ic ImagesCarousel) GetAllImagesCarouselService(ctx context.Context) ([]model.ImageCarousel, error) {
	result, err := ic.Storage.GetAllImagesCarouselStorage(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred in GetAllImagesCarouselStorage: %w", err)
	}

	if len(result) < minimalResult {
		return []model.ImageCarousel{}, model.ErrNoContent
	}

	return result, nil
}

func (ic ImagesCarousel) UpdateImagesCarouselByIDService(ctx context.Context, imagesCarousel model.ImageCarousel) error {
	if err := ic.Storage.UpdateImagesCarouselByIDStorage(ctx, imagesCarousel); err != nil {
		return fmt.Errorf("error occurred in UpdateImagesCarouselByIDStorage: %w", err)
	}

	return nil
}

func (ic ImagesCarousel) DeleteImagesCarouselByIDService(ctx context.Context, param string) error {
	if err := ic.Storage.DeleteImagesCarouselByIDStorage(ctx, param); err != nil {
		return fmt.Errorf("error occurred in DeleteImagesCarouselByIDStorage: %w", err)
	}

	return nil
}
