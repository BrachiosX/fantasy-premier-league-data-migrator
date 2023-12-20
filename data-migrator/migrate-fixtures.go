package datamigrator

import (
	"context"
	"strconv"

	fplapiconnector "github.com/zest97/fpl-player-data-migration/fpl-api-connector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func migrateFixturesData(database *mongo.Database) {
	// migrate match fixtures
	migrateFixtures(database)

	// migrate match histories
	migrateMatchHistories(database)
}

func migrateFixtures(database *mongo.Database) {
	fixtures := fplapiconnector.GetFixtures()

	coll := database.Collection("fixtures")

	for _, fixture := range fixtures {
		filter := bson.D{{"id", fixture.Id}}

		result := coll.FindOneAndReplace(context.Background(), filter, fixture)

		// if no id find to replace, insert
		if result.Err() == mongo.ErrNoDocuments {
			coll.InsertOne(context.Background(), fixture)
		}
	}
}

func migrateMatchHistories(database *mongo.Database) {
	// get all players from collection
	players := getPlayers(database)

	// print number of players
	println("Number of players: " + strconv.Itoa(len(players)))

	// for each player, get match history
	for _, player := range players {
		migrateMatchHistoryForPlayer(database, player)
	}
}

func migrateMatchHistoryForPlayer(database *mongo.Database, player bson.M) {
	// get match history for player
	matchHistory := fplapiconnector.GetMatchHistory(int(player["id"].(int32)))

	// insert match history into database
	coll := database.Collection("match_histories")

	for _, match := range matchHistory.History {
		filter := bson.D{
			{Key: "element", Value: match.Element},
			{Key: "fixture", Value: match.Fixture},
		}

		result := coll.FindOne(context.Background(), filter)

		// if no id find to replace, insert
		if result.Err() == mongo.ErrNoDocuments {
			coll.InsertOne(context.Background(), match)
		}
	}
}

func getPlayers(database *mongo.Database) []bson.M {
	coll := database.Collection("players")

	cursor, err := coll.Find(context.Background(), bson.D{})

	if err != nil {
		panic(err)
	}

	var players []bson.M

	if err = cursor.All(context.Background(), &players); err != nil {
		panic(err)
	}

	return players
}
