package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	configFile string
	local      bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "configs/config.yaml", "The config file.")
	rootCmd.PersistentFlags().BoolVarP(&local, "local", "l", false, "Run on local (true or false)")
}
