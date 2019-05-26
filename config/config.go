package config

import (
	"os"
	"encoding/json"
)

// ConfigFile specifies the JSON file's name
const ConfigFile = "config/config.json"

// Config is about fetching the path from the JSON file
type Config struct {
	Recent struct {
		Title string `json:"title"`
		Data []struct {
			Title string `json:"title"`
		}`json:"data"`
	}`json:"recent"`
	Visit struct {
		Title string `json:"title"`
		Data []struct {
			Title string `json:"title"`
		}`json:"data"`
	} `json:"visit"`
	Trend struct {
		Title string `json:"title"`
		Data []struct{
			Title string `json:"title"`
		} `json:"data"`
	} `json:"trend"`
	Other struct {
		Title string `json:"title"`
		Data []struct {
			Title string `json:"title"`
		}`json:"data"`
	} `json:"other"`
	Headlines struct {
		Data []struct {
			Subtitle string `json:"subtitle"`
			Title string `json:"title"`
			Description string `json:"description"`
			Image string `json:"image"`
		} `json:"data"`
	} `json:"headline"`
}

// LoadConfiguration set the configuration
func LoadConfiguration(path string) (Config, error) {
	var config Config

	dir, err := os.Open(path)
	defer dir.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(dir)
	jsonParser.Decode(&config)
	return config, err
}