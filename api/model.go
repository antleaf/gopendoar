package api

// OpenDOARSystemMetadata is the system metadata for the repository
type OpenDOARSystemMetadata struct {
	ID int `json:"id"`
}

// OpenDOARIdentifier is the identifier for the repository
type OpenDOARIdentifier struct {
	Identifier     string `json:"identifier"`
	IdentifierType string `json:"type"`
}

// OpenDOARName is the name of the repository
type OpenDOARName struct {
	Name     string `json:"name"`
	Acronym  string `json:"acronym"`
	Language string `json:"language"`
}

// OpenDOARSoftware is the software platform that the repository uses
type OpenDOARSoftware struct {
	Name string `json:"name"`
}

// OpenDOAROrganisation is the organisation that the repository belongs to
type OpenDOAROrganisation struct {
	Name        []OpenDOARName       `json:"name"`
	URL         string               `json:"url"`
	Country     string               `json:"country"`
	Identifiers []OpenDOARIdentifier `json:"identifiers"`
}

// OpenDOARRepositoryMetadata is the metadata for the repository
type OpenDOARRepositoryMetadata struct {
	Name            []OpenDOARName   `json:"name"`
	RepoType        string           `json:"type"`
	URL             string           `json:"url"`
	OAIURL          string           `json:"oai_url"`
	Software        OpenDOARSoftware `json:"software"`
	ContentTypes    []string         `json:"content_types"`
	ContentSubjects []string         `json:"content_subjects"`
}

// OpenDOARRepository is the main repository struct
type OpenDOARRepository struct {
	SystemMetadata     OpenDOARSystemMetadata     `json:"system_metadata"`
	Organisation       OpenDOAROrganisation       `json:"organisation"`
	RepositoryMetadata OpenDOARRepositoryMetadata `json:"repository_metadata"`
}
