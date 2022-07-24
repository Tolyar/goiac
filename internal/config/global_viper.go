package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitGlobalViper(cfg *viper.Viper, AppName string, cfgFile string) error {
	if cfgFile != "" {
		// Use config file from the flag.
		cfg.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// name of config file (without extension)
		cfg.SetConfigName(AppName)
		// REQUIRED if the config file does not have the extension in the name
		cfg.SetConfigType("yaml")
		// path to look for the config file in
		cfg.AddConfigPath("/etc")
		// call multiple times to add many search paths
		cfg.AddConfigPath(fmt.Sprintf("%s/.%s", home, AppName))
		cfg.AddConfigPath(".") // optionally look for config in the working directory
	}

	cfg.SetEnvPrefix(AppName)
	cfg.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := cfg.ReadInConfig(); err == nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			return err
		}
	}
	return nil
}
