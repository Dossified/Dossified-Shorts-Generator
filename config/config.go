package config

import (
	"os"

	"github.com/Dominique-Roth/Dossified-Shorts-Generator/utils"

	"github.com/pelletier/go-toml/v2"
)

type ConfigStruct struct {
	LogLevel      string
	RemoteUrl     string
	GowitnessHost string
}

func GetConfiguration() ConfigStruct {
	config := ConfigStruct{}
	config_file_path := "config.toml"
	err := toml.Unmarshal(getConfigFile(config_file_path), &config)
	utils.CheckError(err)
	return config
}

func getConfigFile(file_path string) []byte {
	executable_path, err := os.Getwd()
	utils.CheckError(err)
	content, err := os.ReadFile(executable_path + "/" + file_path)
	utils.CheckError(err)
	return content
}
