package goiac

import (
	"github.com/Tolyar/goiac/internal/config"
	"github.com/Tolyar/goiac/internal/log"
	"github.com/rs/zerolog"
)

// Goiac main object.
var Log *zerolog.Logger

type GoIAC struct {
	Globals config.Globals
	Log     *zerolog.Logger
	Stages  []*Stage
	Scripts []*Script
	Modules []*Module
	Project *Project
}

func NewGoIAC(g config.Globals) *GoIAC {
	goiac := GoIAC{
		Log:     log.InitLog(g.LogLevel),
		Globals: g,
	}

	Log = goiac.Log
	goiac.Stages = make([]*Stage, 0)
	goiac.Scripts = make([]*Script, 0)

	return &goiac
}

func (g *GoIAC) ReadConfig() error {
	if g.Globals.ModulePath != "" {
		module, err := ReadModule(g.Globals.ModulePath, nil, 0)
		if err != nil {
			return err
		}
		g.Scripts = module.S
		g.Modules = append(g.Modules, module)
	}

	if g.Globals.ScriptPath != "" {
		script, err := ReadScript(g.Globals.ScriptPath, nil, 0)
		if err != nil {
			return err
		}
		g.Scripts = append(g.Scripts, script)
	}

	return nil
}
