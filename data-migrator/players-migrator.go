package datamigrator

import (
	"context"
	"fmt"

	"github.com/zest97/fpl-player-data-migration/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func migratePlayers(db *mongo.Database, players []models.Player) {
	coll := db.Collection("players")

	// count players from arg
	fmt.Printf("Prepare migration for total players - %d\n", len(players))

	// filter players
	filteredPlayers := filterPlayers(players, coll)

	// count filtered players
	fmt.Printf("Prepare migration for filtered players - %d\n", len(filteredPlayers))

	for _, player := range filteredPlayers {
		migratePlayer(player, coll)
	}
}

func filterPlayers(players []models.Player, coll *mongo.Collection) []models.Player {
	var filteredPlayers []models.Player

	for _, player := range players {
		// filter players that contain a null value for the "now_cost" field
		// this is a bug in the FPL API
		isNowCostNull := player.NowCost == 0

		// filter players that contain a not null value for the "news" field
		isUnavailable := player.News != ""

		if !isNowCostNull && !isUnavailable {
			filteredPlayers = append(filteredPlayers, player)
		} else {
			// Find and Delete Unavailable Player from system
			filter := bson.D{{"id", player.ID}}
			coll.FindOneAndDelete(context.Background(), filter)
		}
	}

	return filteredPlayers
}

func migratePlayer(player models.Player, coll *mongo.Collection) {
	// get first document and replace it
	filter := bson.D{{"id", player.ID}}
	result := coll.FindOneAndReplace(context.Background(), filter, player)

	// if document is nil, insert it
	if result.Err() == mongo.ErrNoDocuments {
		_, err := coll.InsertOne(context.Background(), player)

		if err != nil {
			panic(err)
		}
	}
}
