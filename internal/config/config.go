package config

import (
	"encoding/base64"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/vnnyx/betty-BE/internal/log"
)

type Configuration struct {
	Environments Environments `mapstructure:"environments"`
}

type Environments struct {
	Development Config `mapstructure:"development"`
	Release     Config `mapstructure:"release"`
	Production  Config `mapstructure:"production"`
}

type Config struct {
	BaseURL       string         `mapstructure:"base-url"`
	Database      DatabaseConfig `mapstructure:"database"`
	RSAPrivateKey string
	RSAPublicKey  string
	EncryptKey    string       `mapstructure:"encrypt-key"`
	AWS           AWSConfig    `mapstructure:"aws"`
	Google        GoogleConfig `mapstructure:"google"`
}

type DatabaseConfig struct {
	Cockroach CockroachConfig `mapstructure:"cockroach"`
}

type GoogleConfig struct {
	ClientID     string `mapstructure:"client-id"`
	ClientSecret string `mapstructure:"client-secret"`
}

type AWSConfig struct {
	AccessKey string `mapstructure:"access-key"`
	SecretKey string `mapstructure:"secret-key"`
	S3URL     string `mapstructure:"s3-url"`
	S3Bucket  string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
}

type CockroachConfig struct {
	DSN      string `mapstructure:"dsn"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

func NewConfig() (*Config, error) {
	logger := log.NewLog()

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal().Caller().Err(err).Msg("Error reading config file")
	}

	var config Configuration

	if err := viper.Unmarshal(&config); err != nil {
		logger.Fatal().Caller().Err(err).Msg("Error unmarshal config file")
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	var currentEnvConfig Config

	switch env {
	case "development":
		currentEnvConfig = config.Environments.Development
	case "release":
		currentEnvConfig = config.Environments.Release
	case "production":
		currentEnvConfig = config.Environments.Production
	default:
		panic("Invalid config")
	}

	privateKey, err := os.ReadFile("./config/private.key")
	if err != nil {
		logger.Fatal().Caller().Err(err).Msg("Error reading private key file")
	}

	publicKey, err := os.ReadFile("./config/public.key")
	if err != nil {
		logger.Fatal().Caller().Err(err).Msg("Error reading public key file")
	}

	currentEnvConfig.RSAPrivateKey = string(privateKey)
	currentEnvConfig.RSAPublicKey = string(publicKey)

	if err := decodeAWSConfig(&currentEnvConfig.AWS, logger); err != nil {
		return nil, err
	}

	if err := decodeGoogleConfig(&currentEnvConfig.Google, logger); err != nil {
		return nil, err
	}

	return &currentEnvConfig, nil
}

func decodeAWSConfig(awsConfig *AWSConfig, logger zerolog.Logger) error {
	decodedAccessKey, err := base64.StdEncoding.DecodeString(awsConfig.AccessKey)
	if err != nil {
		logger.Error().Caller().Err(err).Msg("Error decoding AWS access key")
		return err
	}

	decodedSecretKey, err := base64.StdEncoding.DecodeString(awsConfig.SecretKey)
	if err != nil {
		logger.Error().Caller().Err(err).Msg("Error decoding AWS secret key")
		return err
	}

	awsConfig.AccessKey = string(decodedAccessKey)
	awsConfig.SecretKey = string(decodedSecretKey)

	return nil
}

func decodeGoogleConfig(googleConfig *GoogleConfig, logger zerolog.Logger) error {
	decodedClientID, err := base64.StdEncoding.DecodeString(googleConfig.ClientID)
	if err != nil {
		logger.Error().Caller().Err(err).Msg("Error decoding Google client ID")
		return err
	}

	decodedClientSecret, err := base64.StdEncoding.DecodeString(googleConfig.ClientSecret)
	if err != nil {
		logger.Error().Caller().Err(err).Msg("Error decoding Google client secret")
		return err
	}

	googleConfig.ClientID = string(decodedClientID)
	googleConfig.ClientSecret = string(decodedClientSecret)

	return nil
}
