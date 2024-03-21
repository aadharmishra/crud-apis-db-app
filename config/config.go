package config

import (
	"encoding/json"
	"log"
	"os"
)

var configModel *Config

func NewConfig() (IConfig, error) {
	localConfig, err := initialize()

	if err != nil || localConfig == nil {
		log.Fatalln("config.fetch.error", "Error fetching config locally", map[string]interface{}{"error": err.Error()})
		return nil, err
	}
	return &IConfigModel{model: configModel}, nil
}

// Get implements the interface function for IConfig
func (ic *IConfigModel) Get() *Config {
	return ic.model
}

func initialize() (IConfig, error) {

	var err error
	var data []byte

	configFile := "config.json"

	data, err = os.ReadFile(configFile)
	if err != nil || data == nil {
		return nil, err
	}

	err = json.Unmarshal(data, &configModel)

	if err != nil || configModel == nil {
		log.Fatalf("Error getting creds: %s", err)
		return nil, err
	}

	return &IConfigModel{model: configModel}, nil
}
