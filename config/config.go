package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AppName     string
	Port        string
	Environment string
	Version     string
	DbConfig    DbConfig
	JwtSecret   string
	AWSBucket   string
}

var AppConfig Config

func InitConfig() {
	viper.SetDefault("PORT", "14045")
	viper.SetDefault("ENVIRONMENT", "development")

	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Unable to read config file: %v\n", err)
	}
	checkConfigKey("AWS_ACCESS_KEY_ID", true)
	checkConfigKey("AWS_SECRET_ACCESS_KEY", true)
	checkConfigKey("AWS_BUCKET", true)
	AppConfig = Config{
		AppName:     GetAppName(),
		Port:        GetPort(),
		Environment: GetEnvironment(),
		Version:     GetVersion(),
		DbConfig:    NewDbConfig(),
		AWSBucket:   viper.GetString("AWS_BUCKET"),
	}

}

func GetEnvironment() string {
	checkConfigKey("ENVIRONMENT", true)
	return viper.GetString("ENVIRONMENT")
}
func GetPort() string {
	checkConfigKey("PORT", true)
	return viper.GetString("PORT")
}

func GetAppName() string {
	checkConfigKey("APP_NAME", true)
	return viper.GetString("APP_NAME")
}

func GetVersion() string {
	checkConfigKey("VERSION", false)
	return getStringValueOrDefault("VERSION", "0.0.1")
}

func GetJwtSecret() string {
	checkConfigKey("JWT_SECRET", true)
	return viper.GetString("JWT_SECRET")
}

func checkConfigKey(key string, mandatory bool) {
	if !viper.IsSet(key) {
		if mandatory {
			panic(fmt.Sprintf("Environment variable %s is not set\n", key))
		} else {
			fmt.Printf("Environment variable %s is not set. Using defaults.\n", key)
		}
	}
}

func getStringValueOrDefault(key string, defaultValue string) string {
	if !viper.IsSet(key) {
		return defaultValue
	} else {
		return viper.GetString(key)
	}
}

func getIntegerValueOrDefault(key string, defaultValue int) int {
	if !viper.IsSet(key) {
		return defaultValue
	} else {
		return viper.GetInt(key)
	}
}
