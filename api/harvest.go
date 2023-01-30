package api

import (
	"fmt"
	"os"
	"path/filepath"
)

func Harvest(itemType string, format string) {
	logger.Infof("Harvesting from OpenDOAR into %s....", config.HarvestDataFolderPath)
	client := NewGetClient()
	page := 1
	offset := 0
	for page <= config.HarvestPageLimit {
		logger.Debugf("Harvesting page %d...", page)
		url := fmt.Sprintf("%s?item-type=%s&api-key=%s&order=id&format=%s&limit=%d&offset=%d", config.ApiBaseURL, itemType, os.Getenv(config.ApiKeyENVName), format, config.HarvestPageSize, offset)
		logger.Debugf("Using URL: %s", url)
		status, body, err := APIGet(client, url)
		if err != nil {
			logger.Error(err.Error())
		} else {
			logger.Debugf("Status = %d", status)
			list := APIRepositoryItemList{}
			err = list.UnMarshall(body)
			if err != nil {
				logger.Error(err.Error())
			}
			if list.IsEmpty() == false {
				list.MarshallToFile(filepath.Join(config.HarvestDataFolderPath, fmt.Sprintf("%d.json", page)))
			} else {
				logger.Infof("Reached the end of the results, halting.")
				return
			}
		}
		page += 1
		offset += config.HarvestPageSize
	}
	logger.Infof("Harvest from OpenDOAR into %s completed", config.HarvestDataFolderPath)
}
