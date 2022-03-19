package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

var (
	jwtSecretKey string
)

type Config struct {
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
	//MongoURLKey  string `mapstructure:"MONGO_URL"`
}

const (
	envFile = ".env"
)

func InitializeConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(envFile)
	viper.AutomaticEnv()
	config := Config{}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

func GetJwtSecret() string {
	if jwtSecretKey == "" {
		jwtSecretKey = viper.GetString(JWT_SECRET_KEY)
	}
	return jwtSecretKey
}