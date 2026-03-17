package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppPort     string `mapstructure:"APP_PORT"`
	AppEnv      string `mapstructure:"APP_ENV"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		// If .env is not found, we still want to read from environment variables
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
	}

	err = viper.Unmarshal(&config)
	return
}
