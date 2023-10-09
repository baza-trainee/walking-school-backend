package storage

import (
	"errors"
	"fmt"
	"strings"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	matchedOneDocument      = 1
	contextDeadlineExceeded = "context deadline exceeded"
)

func limitAndOffset(limit, offset int) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))

	return findOptions
}

func creationFilter(phone, email string) primitive.D {
	if phone != "" && email == "" {
		return bson.D{{Key: "phone", Value: phone}}
	} else if email != "" && phone == "" {
		return bson.D{{Key: "email", Value: email}}
	} else {
		return bson.D{
			{Key: "$or", Value: bson.A{
				bson.D{{Key: "phone", Value: phone}},
				bson.D{{Key: "email", Value: email}},
			}},
		}
	}
}

func updateFilter(id, phone, email string) primitive.D {
	if phone != "" && email == "" {
		return bson.D{
			{Key: "$and", Value: bson.A{
				bson.D{{Key: "phone", Value: phone}},
				bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: id}}}},
			}},
		}
	} else if email != "" && phone == "" {
		return bson.D{
			{Key: "$and", Value: bson.A{
				bson.D{{Key: "email", Value: email}},
				bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: id}}}},
			}},
		}
	} else {
		return bson.D{
			{Key: "$or", Value: bson.A{
				bson.D{
					{Key: "$and", Value: bson.A{
						bson.D{{Key: "phone", Value: phone}},
						bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: id}}}},
					}},
				},
				bson.D{
					{Key: "$and", Value: bson.A{
						bson.D{{Key: "email", Value: email}},
						bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: id}}}},
					}},
				},
			}},
		}
	}
}

func handleError(message string, err error) error {
	mce := new(mongo.CommandError)

	if errors.As(err, &mce) && strings.Contains(mce.Message, contextDeadlineExceeded) {
		return model.ErrRequestTimeout
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.ErrNotFound
	}

	return fmt.Errorf("%s: %w", message, err)
}

func handleUpdateByIDError(result *mongo.UpdateResult, message string, err error) error {
	if err != nil {
		handleError(message, err)
	}

	if result.MatchedCount != matchedOneDocument {
		return model.ErrNotFound
	}

	return nil
}

func handleDeleteByIDError(result *mongo.DeleteResult, message string, err error) error {
	if err != nil {
		handleError(message, err)
	}

	if result.DeletedCount != matchedOneDocument {
		return model.ErrNotFound
	}

	return nil
}
