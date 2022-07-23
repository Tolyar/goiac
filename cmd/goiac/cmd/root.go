/*
Copyright © 2022 none

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Tolyar/goiac/internal/config"
	"github.com/Tolyar/goiac/internal/log"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	logLevel     string
	globals      config.Globals
	logLevelsMap = map[string]zerolog.Level{
		"trace":    zerolog.TraceLevel,
		"debug":    zerolog.DebugLevel,
		"info":     zerolog.InfoLevel,
		"warn":     zerolog.WarnLevel,
		"error":    zerolog.ErrorLevel,
		"fatal":    zerolog.FatalLevel,
		"panic":    zerolog.PanicLevel,
		"disabled": zerolog.Disabled,
	}
)

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

	rootCmd.PersistentFlags().StringVarP(&globals.ProjectPath, "project", "p", ".", "path to project directory")
	rootCmd.PersistentFlags().StringVarP(&globals.ModulePath, "module", "m", "", "path to module's directory")
	rootCmd.PersistentFlags().StringVarP(&globals.ScriptPath, "script", "s", "", "path to script")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "loglevel", "l", "info", "Log level: trace, debug, info, warn, error, fatal, panic, disable")
	rootCmd.MarkFlagsMutuallyExclusive("module", "script")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if l, ok := logLevelsMap[logLevel]; ok {
		globals.Log = *log.InitLog(l)
	} else {
		cobra.CheckErr(fmt.Errorf("Loglevel '%v' is incorrect. Possible values is: race, debug, info, warn, error, fatal, panic, disable.", logLevel))
	}
	globals.ReadConfiguration()
}
