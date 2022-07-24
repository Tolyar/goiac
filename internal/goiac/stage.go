package goiac

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type Stage struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description,omitempty"`
	Provider    string `mapstructure:"provider"`
	// Options for provider.
	Options *interface{} `mapstructure:"options,omitempty"`
	Script  *Script
}

func (s *Stage) ValidateStage(script *Script, idx int) error {
	s.Script = script

	// Name is mandatory field.
	if s.Name == "" {
		Log.Error().Str("script", script.Name).Int("idx", idx).Msgf("Name is mandatory filed for stage.")
		Log.Trace().Str("script", script.Name).Msgf("Stage object: %s", pp.Sprint(*s))
		return fmt.Errorf("Script '%v': Name is mandatory filed for stage.", script.Name)
	}

	// Provider is mandatory field.
	if s.Provider == "" {
		Log.Error().Str("script", script.Name).Int("idx", idx).Msgf("Provider is mandatory filed for stage: %s", pp.Sprint(*s))
		Log.Trace().Str("script", script.Name).Msgf("Stage object: %s", pp.Sprint(*s))
		return fmt.Errorf("Script '%v': Provider is mandatory filed for stage.", script.Name)
	}

	return nil
}
