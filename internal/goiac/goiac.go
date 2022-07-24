package goiac

import (
	"github.com/Tolyar/goiac/internal/config"
	"github.com/Tolyar/goiac/internal/log"
	"github.com/rs/zerolog"
)

// Goiac main object.

type GoIAC struct {
	Globals config.Globals
	Log     *zerolog.Logger
	Stages  []Stage
	Scripts []Script
}

func NewGoIAC(g config.Globals) *GoIAC {
	goiac := GoIAC{
		Log:     log.InitLog(g.LogLevel),
		Globals: g,
	}

	goiac.Stages = make([]Stage, 0)
	goiac.Scripts = make([]Script, 0)

	return &goiac
}

func (g *GoIAC) ReadConfig() error {
	var script *Script
	var err error

	if g.Globals.ScriptPath != "" {
		script, err = ReadScript(g.Globals.ScriptPath, g.Log, nil)
		if err != nil {
			return err
		}
	}

	g.Scripts = append(g.Scripts, *script)
	return nil
}
