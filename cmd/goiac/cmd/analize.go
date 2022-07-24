/*
Copyright © 2022 none

*/
package cmd

import (
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

// analizeCmd represents the analize command
var analizeCmd = &cobra.Command{
	Use:   "analize",
	Short: "Analize configuration file.",
	Long:  `Analize configuration for debug purposes.`,
	Run: func(cmd *cobra.Command, args []string) {
		G.Log.Info().Str("object", "globals").Msg(pp.Sprint(G.Globals))
		G.Log.Info().Str("object", "facts").Msg(pp.Sprint(G.Facts))
		for i, m := range G.Modules {
			G.Log.Info().Str("object", "module").Int("idx", i).Msg(pp.Sprint(m))
			for j, s := range m.S {
				G.Log.Debug().Str("object", "script").Int("idx", j).Str("module", m.Name).Msg(pp.Sprint(s))
			}
		}
		for i, s := range G.Scripts {
			G.Log.Info().Str("object", "script").Int("idx", i).Msg(pp.Sprint(s))
			for j, st := range s.Stages {
				G.Log.Debug().Str("object", "stage").Int("idx", j).Str("script", s.Name).Msg(pp.Sprint(st))
			}
		}
		if G.Project != nil {
			G.Log.Info().Str("object", "project").Msg(pp.Sprint(G.Project))
		}
	},
}

func init() {
	rootCmd.AddCommand(analizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// analizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// analizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
