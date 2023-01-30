package api

type OpenDOARSystemMetadata struct {
	ID int `json:"id"`
}

type OpenDOARIdentifier struct {
	Identifier     string `json:"identifier"`
	IdentifierType string `json:"type"`
}

type OpenDOARName struct {
	Name     string `json:"name"`
	Acronym  string `json:"acronym"`
	Language string `json:"language"`
}

type OpenDOARSoftware struct {
	Name string `json:"name"`
}

type OpenDOAROrganisation struct {
	Name        []OpenDOARName       `json:"name"`
	URL         string               `json:"url"`
	Country     string               `json:"country"`
	Identifiers []OpenDOARIdentifier `json:"identifiers"`
}

type OpenDOARRepositoryMetadata struct {
	Name            []OpenDOARName   `json:"name"`
	RepoType        string           `json:"type"`
	URL             string           `json:"url"`
	OAIURL          string           `json:"oai_url"`
	Software        OpenDOARSoftware `json:"software"`
	ContentTypes    []string         `json:"content_types"`
	ContentSubjects []string         `json:"content_subjects"`
}

type OpenDOARRepository struct {
	SystemMetadata     OpenDOARSystemMetadata     `json:"system_metadata"`
	Organisation       OpenDOAROrganisation       `json:"organisation"`
	RepositoryMetadata OpenDOARRepositoryMetadata `json:"repository_metadata"`
}
