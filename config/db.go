package config

import "github.com/spf13/viper"

type DbConfig struct {
	DatabaseType string
	DatabaseLink string
	DatabasePort string
	DatabaseName string
	DatabaseUser string
	DatabasePass string
}

func NewDbConfig(dbType string, dbLink string) DbConfig {
	checkConfigKey("DB_TYPE", true)
	checkConfigKey("DB_LINK", true)
	checkConfigKey("DB_PORT", true)
	if viper.GetString("DB_TYPE") != "sqlite" {
		checkConfigKey("DB_USER", true)
		checkConfigKey("DB_PASS", true)
		checkConfigKey("DB_NAME", true)
	}
	return DbConfig{
		DatabaseType: getStringValueOrDefault("DB_TYPE", "sqlite"),
		DatabaseLink: getStringValueOrDefault("DB_LINK", "mitsukeru.db"),
		DatabasePort: getStringValueOrDefault("DB_PORT", "5432"),
		DatabaseName: getStringValueOrDefault("DB_NAME", "mitsukeru-server"),
		DatabaseUser: viper.GetString("DB_USER"),
		DatabasePass: viper.GetString("DB_PASS"),
	}
}
