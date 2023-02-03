package cmd

import (
	"fmt"
	"os"
)

// InitialiseApplication initialises the application
func InitialiseApplication() {
	err := (&config).Initialise(ConfigPath)
	if err != nil {
		fmt.Print(err.Error() + "\n")
		fmt.Printf("Halting execution because config file not loaded from %s\n", ConfigPath)
		os.Exit(1)
	}
	logger, err = ConfigureZapSugarLogger(Debug)
	if Debug {
		logger.Infof("Debugging enabled")
	}
	if err != nil {
		logger.Fatal("Unable to initialise logging, halting")
		os.Exit(-1)
	}
}
