package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment         string        `mapstructure:"ENVIRONMENT"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	AllowHeaders        []string      `mapstructure:"ALLOW_HEADERS"`
	RedisAddress        string        `mapstructure:"REDIS_ADDRESS"`
	RedisPassword       string        `mapstructure:"REDIS_PASSWORD"`
	MigrationURL        string        `mapstructure:"MIGRATION_URL"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	NotificationsPort   string        `mapstructure:"NOTIFICATIONS_PORT"`
	OrderPort           string        `mapstructure:"ORDER_PORT"`
	PaymentPort         string        `mapstructure:"PAYMENT_PORT"`
	ProductPort         string        `mapstructure:"PRODUCT_PORT"`
	KeySeed             string        `mapstructure:"KEY_SEED"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
