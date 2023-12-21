package datamigrator

import (
	"context"

	"github.com/BrachiosX/fantasy-premier-league-data-migrator/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func migrateGameSettings(db *mongo.Database, gameSettings models.Setting) {
	coll := db.Collection("game_settings")

	// get first document and replace it
	result := coll.FindOneAndReplace(context.Background(), bson.D{{}}, gameSettings)

	// if document is nil, insert it
	if result.Err() == mongo.ErrNoDocuments {
		_, err := coll.InsertOne(context.Background(), gameSettings)

		if err != nil {
			panic(err)
		}
	}
}
