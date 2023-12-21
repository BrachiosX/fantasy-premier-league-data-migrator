package datamigrator

import (
	"context"

	"github.com/BrachiosX/fantasy-premier-league-data-migrator/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func migrateTeams(db *mongo.Database, teams []models.Team) {
	coll := db.Collection("teams")

	// get first document and replace it
	for _, team := range teams {
		println(team.Name)

		filter := bson.D{{"id", team.ID}}
		result := coll.FindOneAndReplace(context.Background(), filter, team)

		// if document is nil, insert it
		if result.Err() == mongo.ErrNoDocuments {
			_, err := coll.InsertOne(context.Background(), team)

			if err != nil {
				panic(err)
			}
		} else if result.Err() != nil {
			panic(result.Err())
		}
	}
}
