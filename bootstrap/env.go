package Bootstrap

import (
	"errors"
	"fmt"
	"log"
	"os"

	Loghelper "testing-api/loghelper"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	ExternalApiBaseUrl     string `mapstructure:"EXTERNAL_API_BASE_URL"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		// in debug mode
		viper.SetConfigFile("../.env")
	}

	err := viper.ReadInConfig()
	if err != nil {
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Can't find the file .env: %s", err))
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Environment can't be loaded: %s", err))
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
		Loghelper.WriteLog().Info().Msg("The App is running in development env")
	}

	return &env
}
