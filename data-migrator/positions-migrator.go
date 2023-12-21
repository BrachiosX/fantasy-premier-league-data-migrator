package datamigrator

import (
	"context"

	"github.com/BrachiosX/fantasy-premier-league-data-migrator/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func migratePositions(db *mongo.Database, positions []models.Position) {
	coll := db.Collection("positions")

	for _, position := range positions {
		// get first document and replace it
		result := coll.FindOneAndReplace(context.Background(), bson.D{{"id", position.ID}}, position)

		// if document is nil, insert it
		if result.Err() == mongo.ErrNoDocuments {
			_, err := coll.InsertOne(context.Background(), position)

			if err != nil {
				panic(err)
			}
		}
	}
}
