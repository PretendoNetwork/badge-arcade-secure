package database

import (
	"context"

	"github.com/PretendoNetwork/nintendo-badge-arcade-secure/globals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DoesSessionExist(pid uint32) bool {
	var result bson.M

	err := sessionsCollection.FindOne(context.TODO(), bson.D{{"pid", pid}}, options.FindOne()).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		} else {
			globals.Logger.Error(err.Error())
			return false
		}
	} else {
		return true
	}
}
