package datamigrator

import (
	"context"

	"github.com/zest97/fpl-player-data-migration/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func migratePlayerStats(db *mongo.Database, playerStats []models.PlayerStatTypes) {
	coll := db.Collection("stats")

	for _, statData := range playerStats {
		filter := bson.D{{Key: "name", Value: statData.Label}}

		// get first document and replace it
		result := coll.FindOneAndReplace(context.Background(), filter, statData)

		// if document is nil, insert it
		if result.Err() == mongo.ErrNoDocuments {
			_, err := coll.InsertOne(context.Background(), statData)

			if err != nil {
				panic(err)
			}
		}
	}
}
