package providers

import (
	"encoding/json"

	"github.com/crhuber/cellar/pkg/core"
)

type CellarExport struct {
	Version   string                   `json:"version"`
	Providers map[string]core.MetaInfo `json:"providers"`
}

func GenerateProvidersMetaJSON(version string, providersMetaList []core.MetaInfo) (string, error) {
	providersMetaMap := make(map[string]core.MetaInfo)
	for _, provider := range providersMetaList {
		providersMetaMap[provider.Name] = provider
	}

	cellarObject := CellarExport{
		Version:   version,
		Providers: providersMetaMap,
	}

	jsonOutput, err := json.MarshalIndent(cellarObject, "", "  ")

	if err != nil {
		return "", err
	}

	return string(jsonOutput), nil
}
