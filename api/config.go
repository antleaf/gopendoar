package api

// APIConfiguration is the configuration for the API package
type APIConfiguration struct {
	HarvestDataFolderPath string `yaml:"harvestDataFolderPath"`
	HarvestPageSize       int    `yaml:"harvestPageSize"`
	HarvestPageLimit      int    `yaml:"harvestPageLimit"`
	ApiBaseURL            string `yaml:"apiBaseURL"`
	ApiKeyENVName         string `yaml:"apiKeyENVKey"`
}
