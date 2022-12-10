package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBUri  string   `mapstructure:"MONGODB_LOCAL_URI"`
	Port   string   `mapstructure:"PORT"`
	Token  string   `mapstructure:"TOKEN"`
	Admins []string `mapstructure:"ADMINS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.BindEnv("MONGODB_LOCAL_URI")
	viper.BindEnv("PORT")
	viper.BindEnv("TOKEN")
	viper.BindEnv("ADMINS")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	err = viper.Unmarshal(&config)
	log.Println(config.DBUri)
	log.Println(config.Port)
	log.Println(config.Admins)
	return
}
