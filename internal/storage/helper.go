package storage

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	matchedOneDocument = 1
)

func LimitAndOffset(limit, offset int) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))

	return findOptions
}

func contactFilter(phone, email string) primitive.D {
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
