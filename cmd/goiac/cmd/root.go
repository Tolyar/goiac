/*
Copyright © 2022 none

*/
package cmd

import (
	"os"

	"github.com/Tolyar/goiac/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	globals config.Globals
	cfgFile string
	goiac   config.GoIAC
)

const AppName = "goiac"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goiac",
	Short: "Simple configuration manager",
	Long: `Simple IAC software Golang based.
	It does not plan to replace ansible, puppet, and so on.
	GoIAC is only for very simple cases.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			os.Exit(0)
		}
		os.Exit(0)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Path to config file")
	rootCmd.PersistentFlags().StringP("project_path", "p", ".", "path to project directory")
	rootCmd.PersistentFlags().StringP("module_path", "m", "", "path to module's directory")
	rootCmd.PersistentFlags().StringP("script_path", "s", "", "path to script")
	rootCmd.PersistentFlags().StringP("log_level", "l", "", "Log level: trace, debug, info, warn, error, fatal, panic, disable")
	rootCmd.MarkFlagsMutuallyExclusive("module_path", "script_path")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	cfg := viper.New()
	cfg.SetDefault("log_level", "info")
	cfg.SetDefault("project_path", ".")
	cfg.SetDefault("module_path", "")
	cfg.SetDefault("script_path", "")
	err := config.InitGlobalViper(cfg, AppName, cfgFile)
	cobra.CheckErr(err)

	flags := rootCmd.Flags()

	// Read global values from viper and flags.
	globals.LogLevel = config.FlagAndViper("log_level", flags, cfg)
	globals.ProjectPath = config.FlagAndViper("project_path", flags, cfg)
	globals.ModulePath = config.FlagAndViper("module_path", flags, cfg)
	globals.ScriptPath = config.FlagAndViper("script_path", flags, cfg)
}
