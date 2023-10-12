package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/baza-trainee/walking-school-backend/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	connectTimeout           = 5 * time.Second
	projectCollection        = "Project"
	userCollection           = "User"
	partnerCollection        = "Partner"
	contactCollection        = "Contact"
	heroCollection           = "Hero"
	projSectDescCollection   = "ProjectsSectionDescription"
	imagesCarouselCollection = "ImagesCarousel"
	walkingSchoolDatabase    = "WalkingSchool"
	feedbackCollection       = "Feedback"
	connectionString         = "mongodb+srv://%s:%s@%s.fqu5xkh.mongodb.net/"
)

type Storage struct {
	DB *mongo.Database
}

func NewStorage(cfg config.MongoDB) (Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	opt := options.Client().SetConnectTimeout(connectTimeout).ApplyURI(
		fmt.Sprintf(connectionString, cfg.User, cfg.Password, cfg.Host))

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return Storage{}, fmt.Errorf("cannot create Connect: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return Storage{}, fmt.Errorf("error during Ping to db: %w", err)
	}

	database := client.Database(cfg.DbName)

	return Storage{DB: database}, nil
}
