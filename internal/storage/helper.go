package storage

import "go.mongodb.org/mongo-driver/mongo/options"

const (
	matchedOneDocument = 1
)

func LimitAndOffset(limit, offset int) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))

	return findOptions
}
