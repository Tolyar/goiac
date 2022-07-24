package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func FlagAndViper(name string, flags *pflag.FlagSet, cfg *viper.Viper) string {
	f := flags.Lookup(name)
	if f != nil && !f.Changed && cfg.IsSet(f.Name) {
		return cfg.GetString(f.Name)
	}

	return f.Value.String()
}
