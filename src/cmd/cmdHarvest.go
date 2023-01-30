package cmd

import (
	"github.com/spf13/cobra"
	api "gopendoar/src/api"
)

var harvestCmd = &cobra.Command{
	Use: "harvest",
	Run: func(cmd *cobra.Command, args []string) {
		initialiseApplication()
		api.InitialiseAPI(config.APIConfig, logger)
		err := PromptUser("This will over-write the harvested data. Are you sure you want to proceed?")
		if err != nil {
			logger.Info("Aborted.")
			return
		}
		api.Harvest("repository", "Json")
	},
}
