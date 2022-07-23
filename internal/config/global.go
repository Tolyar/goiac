package config

import "github.com/rs/zerolog"

// Global configuration structs.

type Globals struct {
	ProjectPath string
	ModulePath  string
	ScriptPath  string
	Log         zerolog.Logger
}

func (g *Globals) ReadConfiguration() {
}
