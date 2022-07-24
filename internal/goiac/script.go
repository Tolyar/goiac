package goiac

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Script struct {
	Name        string  `mapstructure:"name"`
	Description string  `mapstructure:"description,omitempty"`
	Stages      []Stage `mapstructure:"stages"`
	Module      *string
}

// Read script from file and return Script object.
func ReadScript(path string, log *zerolog.Logger, module *string) (*Script, error) {
	cfg := viper.New()
	cfg.SetConfigFile(path)
	cfg.SetConfigType("yaml")
	err := cfg.ReadInConfig()
	if err != nil {
		log.Error().Err(err).Str("path", path).Msg("Can't read script file.")
		return nil, err
	}
	s := Script{}
	err = cfg.Unmarshal(&s)
	if err != nil {
		log.Error().Err(err).Str("path", path).Msg("Can't parse script file.")
		return nil, err
	}
	s.Module = module

	// Name is mandatory field.
	if s.Name == "" {
		log.Error().Str("path", path).Msg("Name is mandatory filed for scripts")
		return nil, fmt.Errorf("Name is mandatory filed for scripts '%v'", path)
	}

	// Stages is mandatory field.
	if s.Stages == nil {
		log.Error().Str("path", path).Msg("Stages is mandatory filed for scripts")
		return nil, fmt.Errorf("Stages is mandatory filed for scripts '%v'", path)
	}

	return &s, nil
}