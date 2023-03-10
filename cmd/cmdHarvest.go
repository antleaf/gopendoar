package cmd

import (
	"github.com/antleaf/gopendoar/api"
	"github.com/spf13/cobra"
)

// harvestCmd represents the harvest command
var harvestCmd = &cobra.Command{
	Use: "harvest",
	Run: func(cmd *cobra.Command, args []string) {
		InitialiseApplication()
		api.InitialiseAPI(config.APIConfig, logger)
		err := PromptUser("This will over-write the harvested data. Are you sure you want to proceed?")
		if err != nil {
			logger.Info("Aborted.")
			return
		}
		api.Harvest("repository", "Json")
	},
}
