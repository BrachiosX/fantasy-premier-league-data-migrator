package datamigrator

import (
	fplapiconnector "github.com/zest97/fpl-player-data-migration/fpl-api-connector"
	"go.mongodb.org/mongo-driver/mongo"
)

func migrateBootstrapData(database *mongo.Database) {
	// get data from API
	fplBootstrapData := fplapiconnector.GetFPLBoostrapData()

	// migrate game settings
	migrateGameSettings(database, fplBootstrapData.GameSettings)

	// migrate teams
	migrateTeams(database, fplBootstrapData.Teams)

	// migrate players
	migratePlayers(database, fplBootstrapData.Players)

	// migrate positions
	migratePositions(database, fplBootstrapData.Positions)

	// migrate stats
	migratePlayerStats(database, fplBootstrapData.Stats)
}
