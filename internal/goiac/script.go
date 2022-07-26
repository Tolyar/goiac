package goiac

import (
	"fmt"
)

type Script struct {
	Name        string   `mapstructure:"name"`
	Description string   `mapstructure:"description,omitempty"`
	Stages      []*Stage `mapstructure:"stages"`
	Module      *Module
	Idx         int
}

// Read script from file and return Script object.
func ReadScript(path string, module *Module, idx int) (*Script, error) {
	cfg, err := ReadAndTemplate(path)
	if err != nil {
		Log.Error().Err(err).Str("path", path).Msg("Can't read script file.")
		return nil, err
	}
	s := Script{}
	err = cfg.Unmarshal(&s)
	if err != nil {
		Log.Error().Err(err).Str("path", path).Msg("Can't parse script file.")
		return nil, err
	}

	if module != nil {
		s.Module = module
	} else {
		m := Module{
			Name: "w/o module",
		}
		s.Module = &m
	}
	// Name is mandatory field.
	if s.Name == "" {
		Log.Error().Str("path", path).Int("idx", idx).Str("module", s.Module.Name).Msg("Name is mandatory filed for scripts")
		return nil, fmt.Errorf("Name is mandatory filed for scripts '%v'", path)
	}

	// Stages is mandatory field.
	if s.Stages == nil {
		Log.Error().Str("path", path).Int("idx", idx).Str("module", s.Module.Name).Msg("Stages is mandatory filed for scripts")
		return nil, fmt.Errorf("Stages is mandatory filed for scripts '%v'", path)
	}
	s.Idx = idx
	for i, st := range s.Stages {
		if err := st.ValidateStage(&s, i); err != nil {
			return nil, err
		}
	}

	return &s, nil
}
