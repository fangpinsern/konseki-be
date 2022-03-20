package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
	KONSEKI_LINK = "KONSEKI_LINK"
)

var (
	jwtSecretKey string
	konsekiLink string
)

type Config struct {
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
	KonsekiLink string `mapstructure:"KONSEKI_LINK"`
}

const (
	envFile = "./env/.env"
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

func GetKonsekiLink() string {
	if konsekiLink == "" {
		konsekiLink = viper.GetString(KONSEKI_LINK)
	}
	return konsekiLink
}