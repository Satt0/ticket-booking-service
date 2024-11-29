package env

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DB_URL     string `mapstructure:"DB_URL"`
	APP_PORT   string `mapstructure:"APP_PORT"`
	APP_ENV    string `mapstructure:"APP_ENV"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
}

func NewEnv() *Env {

	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}
	return &env
}
