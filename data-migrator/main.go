package datamigrator

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func Migrate(database *mongo.Database) {
	migrateBootstrapData(database)

	migrateFixturesData(database)
}
