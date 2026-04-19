/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "i3dual-pc",
	Short: "This application allow you to jump between two computers running i3wm",
	Long: `i3dual-pc is a CLI application written in Go, having the propouse to 
	improve the expirience of using 2 computers running i3wm.

	The aplication is a very small agent running on the Server machine, exposing 
	a port (custom or default) on the internal network, and one endpoint to be 
	acessed when the user wants to jump between two computers.

	The computer Cliente makes a Curl request to the endpoint exposed by the
	Server computer. Internaly the program runs a few commands to move the mouse 
	into the monitor of the Server/Cliente computers.

	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.i3dual-pc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
