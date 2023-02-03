package api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// RepositoryList is a list of OpenDOAR repository records
type RepositoryList struct {
	Items []OpenDOARRepository `yaml:"items"`
}

// Harvest harvests data from OpenDOAR, writing the results to files in the harvest data folder
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
			list := RepositoryList{}
			err = list.UnMarshall(body)
			if err != nil {
				logger.Error(err.Error())
			}
			if list.IsEmpty() == false {
				filePath := filepath.Join(config.HarvestDataFolderPath, fmt.Sprintf("%d.json", page))
				logger.Debugf("Writing file... %s", filePath)
				err = list.MarshallToFile(filePath)
				if err != nil {
					logger.Error(err.Error())
				}
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

// UnMarshall unmarshalls JSON data in a byte array into a RepositoryList
func (l *RepositoryList) UnMarshall(bytes []byte) error {
	err := json.Unmarshal(bytes, l)
	return err
}

// MarshallToFile marshalls a RepositoryList into a JSON file
func (l *RepositoryList) MarshallToFile(filePath string) error {
	out, err := json.Marshal(l)
	if err != nil {
		return err
	}
	var file *os.File
	file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(out)
	return err
}

// IsEmpty returns true if the RepositoryList is empty
func (l *RepositoryList) IsEmpty() bool {
	return len(l.Items) == 0
}
