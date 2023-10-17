/*
Copyright © 2023 Alan Ton (alandgton@gmail.com)
*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "resy-sniper",
	Short: "CLI to snipe reservations from resy",
	Long: `Resy Sniper is a tool that devs can use to secure the most
coveted restaurant reservations.

The initial plan is to have users provide their login credentials
(e.g. API key) and have this script run a loop to continuously
send booking requests to Resy's servers.`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.resy-sniper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


