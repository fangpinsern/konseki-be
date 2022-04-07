package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

const (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
	KONSEKI_LINK = "KONSEKI_LINK"
	FIRESTORE_TYPE = "FIRESTORE_TYPE"
	FIRESTORE_PROJECT_ID = "FIRESTORE_PROJECT_ID"
	FIRESTORE_PRIVATE_KEY_ID = "FIRESTORE_PRIVATE_KEY_ID"
	FIRESTORE_PRIVATE_KEY = "FIRESTORE_PRIVATE_KEY"
	FIRESTORE_CLIENT_EMAIL = "FIRESTORE_CLIENT_EMAIL"
	FIRESTORE_CLIENT_ID="FIRESTORE_CLIENT_ID"
	FIRESTORE_AUTH_URI= "FIRESTORE_AUTH_URI"
	FIRESTORE_TOKEN_URI = "FIRESTORE_TOKEN_URI"
	FIRESTORE_AUTH_PROVIDER_CERT_URL = "FIRESTORE_AUTH_PROVIDER_CERT_URL"
	FIRESTORE_CLIENT_CERT_URL = "FIRESTORE_CLIENT_CERT_URL"
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

func GetFirestoreCreds() []byte {
	privateKey := viper.GetString(FIRESTORE_PRIVATE_KEY)
	stuff := map[string]string{
		"type": viper.GetString(FIRESTORE_TYPE),
		"project_id": viper.GetString(FIRESTORE_PROJECT_ID),
		"private_key_id": viper.GetString(FIRESTORE_PRIVATE_KEY_ID),
		"private_key": strings.Replace(privateKey, "\\n", "\n", -1),
		"client_email": viper.GetString(FIRESTORE_CLIENT_EMAIL),
		"client_id": viper.GetString(FIRESTORE_CLIENT_ID),
		"auth_uri": viper.GetString(FIRESTORE_AUTH_URI),
		"token_uri": viper.GetString(FIRESTORE_TOKEN_URI),
		"auth_provider_x509_cert_url": viper.GetString(FIRESTORE_AUTH_PROVIDER_CERT_URL),
		"client_x509_cert_url": viper.GetString(FIRESTORE_CLIENT_CERT_URL),
	}
	creds, err := json.Marshal(stuff)

	if err != nil{
		panic("help")
	}

	return creds
}



