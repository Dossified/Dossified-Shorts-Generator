// Helper to configure application according to provided TOML file
package config

import (
	"os"

	"github.com/Dossified/Dossified-Shorts-Generator/utils"

	"github.com/pelletier/go-toml/v2"
)

// Configuration file structure
type ConfigStruct struct {
	LogLevel      string
	RemoteUrl     string
	GowitnessHost string

	AmountNewsTrends     int
	AmountUpcomingEvents int
	AmountDaysTrends     int

	UploadToYouTube   bool
	UploadToInstagram bool

	InstagramUsername string
	InstagramPassword string
}

// Retrieves current configuration from file
func GetConfiguration() ConfigStruct {
	config := ConfigStruct{}
	config_file_path := "config.toml"
	err := toml.Unmarshal(getConfigFile(config_file_path), &config)
	utils.CheckError(err)
	return config
}

// Retrieves content of configuration file
func getConfigFile(file_path string) []byte {
	executable_path, err := os.Getwd()
	utils.CheckError(err)
	content, err := os.ReadFile(executable_path + "/" + file_path)
	utils.CheckError(err)
	return content
}
