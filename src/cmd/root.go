package cmd

import (
	"go.uber.org/zap"
	"os"

	"github.com/spf13/cobra"
)

var logger *zap.SugaredLogger
var config = Configuration{}
var Debug bool
var ConfigPath string

var rootCmd = &cobra.Command{
	Use:   "gopendoar",
	Short: "A simple harvest client for OpenDOAR",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "", false, "--debug=true|false")
	rootCmd.PersistentFlags().StringVarP(&ConfigPath, "config", "", "./config/config.yaml", "--config=<file_path>")
	rootCmd.AddCommand(harvestCmd)
}
