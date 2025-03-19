package modules

import (
  "encoding/json"
  "github.com/google/uuid"
)

// ModuleMetadata represents metadata for a scraping module.
type ModuleMetadata struct {
  SourceName      string `json:"sourceName"`
  Author          Author `json:"author"`
  IconURL         string `json:"iconUrl"`
  Version         string `json:"version"`
  Language        string `json:"language"`
  BaseURL         string `json:"baseUrl"`
  StreamType      string `json:"streamType"`
	Quality         string `json:"quality"`
	SearchBaseURL   string `json:"searchBaseUrl"`
	ScriptURL       string `json:"scriptUrl"`
	AsyncJS         *bool  `json:"asyncJS,omitempty"`
	StreamAsyncJS   *bool  `json:"streamAsyncJS,omitempty"`
	Softsub         *bool  `json:"softsub,omitempty"`
}

// Author represents the author of a module.
type Author struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// ScrapingModule represents a module with metadata and file paths.
type ScrapingModule struct {
	ID           uuid.UUID      `json:"id"`
	Metadata     ModuleMetadata `json:"metadata"`
	LocalPath    string         `json:"localPath"`
	MetadataURL  string         `json:"metadataUrl"`
	IsActive     bool           `json:"isActive"`
}

// NewScrapingModule initializes a new ScrapingModule.
func NewScrapingModule(metadata ModuleMetadata, localPath, metadataURL string, isActive bool) ScrapingModule {
	return ScrapingModule{
		ID:          uuid.New(),
		Metadata:    metadata,
		LocalPath:   localPath,
		MetadataURL: metadataURL,
		IsActive:    isActive,
	}
}

// Serialize converts the ScrapingModule to a JSON string.
func (m ScrapingModule) Serialize() (string, error) {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// DeserializeScrapingModule parses a JSON string into a ScrapingModule.
func DeserializeScrapingModule(jsonData string) (ScrapingModule, error) {
	var module ScrapingModule
	err := json.Unmarshal([]byte(jsonData), &module)
	return module, err
}
