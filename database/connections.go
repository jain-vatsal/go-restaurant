package database

import (
	"github.com/jain-vatsal/go-restaurant/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(collectionName string) *mongo.Collection {
	switch collectionName {
	case constants.FOOD:
		// return OpenCollection(client, constants.FOOD)
	}
	return nil
}
