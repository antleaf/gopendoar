package api

import (
	"encoding/json"
	"os"
)

type APIRepositoryItemList struct {
	Items []OpenDOARRepository `yaml:"items"`
}

func (l *APIRepositoryItemList) UnMarshall(bytes []byte) error {
	err := json.Unmarshal(bytes, l)
	return err
}

func (l *APIRepositoryItemList) MarshallToFile(filePath string) error {
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

func (l *APIRepositoryItemList) IsEmpty() bool {
	return len(l.Items) == 0
}
