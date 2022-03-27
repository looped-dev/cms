package db

import (
	"github.com/looped-dev/cms/api/constants"
	"github.com/spf13/viper"
)

// GetDatabaseName determine the database name to use based on whether the
// config has been set or not. If the config DB name isn't set, it defaults to
// the default database name
func GetDatabaseName() string {
	dbName := viper.GetString(constants.VIPER_DB_NAME)
	if dbName == "" {
		return constants.DefaultDBName
	}
	return dbName
}
