package app

import (
	"github.com/spf13/viper"
	"org.idev.bunny/backend/api/enum"
	"org.idev.bunny/backend/common/logger"
)

type appConfig struct {
	Env         enum.Env `mapstructure:"ENV"`
	Port        string   `mapstructure:"PORT"`
	DbUrl       string   `mapstructure:"DB_URL"`
	JWKsUrl     string   `mapstructure:"JWKS_URL"`
	RedisUrl    string   `mapstructure:"REDIS_URL"`
	KafkaHost   string   `mapstructure:"KAFKA_HOST"`
	KafkaPort   int32    `mapstructure:"KAFKA_PORT"`
	MongoUrl    string   `mapstructure:"MONGO_URL"`
	MongoDbName string   `mapstructure:"MONGO_DB_NAME"`
}

func LoadConfig() (*appConfig, error) {

	log := logger.New("Server", "Load Config")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AllowEmptyEnv(false)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warnf("config not found: %v", err)
		}
		return nil, err
	}

	appConfig := &appConfig{}
	if err := viper.Unmarshal(appConfig); err != nil {
		return nil, err
	}

	return appConfig, nil
}
