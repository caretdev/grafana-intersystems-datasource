package models

import (
	"encoding/json"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

// Settings represents the Datasource options in Grafana
type Settings struct {
	Addr      string `json:"url"`
	Namespace string `json:"database"`
	User      string `json:"user"`
	Password  string `json:"password"`
}

// LoadSettings converts the DataSourceInLoadSettings to usable Github settings
func LoadSettings(settings backend.DataSourceInstanceSettings) (Settings, error) {
	s := Settings{}
	if err := json.Unmarshal(settings.JSONData, &s); err != nil {
		return Settings{}, err
	}
	s.Addr = settings.URL
	s.Namespace = settings.Database
	s.User = settings.User

	if val, ok := settings.DecryptedSecureJSONData["password"]; ok {
		s.Password = val
	}

	return s, nil
}
