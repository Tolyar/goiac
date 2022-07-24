package config

import (
	"github.com/Tolyar/goiac/internal/log"
	"github.com/rs/zerolog"
)

// Goiac main object.

type GoIAC struct {
	Globals Globals
	Log     *zerolog.Logger
}

func NewGoIAC(g Globals) *GoIAC {
	goiac := GoIAC{
		Log: log.InitLog(g.LogLevel),
	}

	return &goiac
}
