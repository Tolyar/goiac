/*
Copyright © 2022 none

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// analizeCmd represents the analize command
var analizeCmd = &cobra.Command{
	Use:   "analize",
	Short: "Analize configuration file.",
	Long:  `Analize configuration for debug purposes.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("analize called")
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
