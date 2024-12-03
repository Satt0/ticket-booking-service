package env

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	DB_URL       string `mapstructure:"DB_URL"`
	APP_PORT     string `mapstructure:"APP_PORT"`
	APP_ENV      string `mapstructure:"APP_ENV"`
	JWT_SECRET   string `mapstructure:"JWT_SECRET"`
	KAFKA_BROKER string `mapstructure:"KAFKA_BROKER"`
}

func NewEnv() *Env {
	v := viper.New()
	env := Env{}
	v.SetConfigFile(".env")
	// load from os env
	v.SetDefault("APP_PORT", os.Getenv("APP_PORT"))
	v.SetDefault("DB_URL", os.Getenv("DB_URL"))
	if os.Getenv("KAFKA_BROKER") != "" {
		v.Set("KAFKA_BROKER", os.Getenv("KAFKA_BROKER"))
	}
	// Automatically search for environment variables
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = v.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}
	return &env
}
