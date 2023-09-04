package common

import "github.com/rs/zerolog"

type Option struct {
	Config *Config
	Log    *zerolog.Logger
}
