package common

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Server struct {
	Host string
	Port string
}
type Log struct {
	Level string
}
type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type Config struct {
	Server Server
	DB     DB
	Log    Log
}

func NewConfig() (conf *Config, err error) {

	v := viper.New()
	v.AddConfigPath("./deployment")
	v.SetConfigName("product-cfg")
	v.SetConfigType("yaml")

	err = v.ReadInConfig()
	if err != nil {
		log.Err(err).Msg("config file is not found")
		return
	} else {
		log.Info().Msgf("using config file %s", v.ConfigFileUsed())
	}

	err = v.Unmarshal(&conf)
	if err != nil {
		return
	}

	return
}
