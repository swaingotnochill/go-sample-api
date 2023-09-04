package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("couldn't find .env file due to : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("environment variables couldn't be loaded due to : ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development mode.")
	}

	return &env
}
