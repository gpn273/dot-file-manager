package main

import (
	"os"
	"encoding/json"
	"path"
)

type GitSettings struct {
	Remote string `json:"remote"`
	Branch string `json:"branch"`
	User string `json:"user"`
	PrivateKey string `json:"private_key"`
}

type ConfigSettings struct {
	Git GitSettings `json:"git"`
	Backup bool `json:"backup"`
}

func ConfigGetFilePath() string  {
	var configFile string = FLAG_CONFIG_FILE
	empty := isEmpty(configFile)

	if empty {
		configFile = path.Join(CONFIG_DEFAULT_FILE_LOCATION, CONFIG_DEFAULT_FILE_NAME)
	}

	return configFile
}

func ConfigExists() bool {
	return pathExists(ConfigGetFilePath())
}

func ParseConfig() {
	configFile, err := os.Open(ConfigGetFilePath())
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to open config file: " + err.Error(),
			Severity: "Error",
			Terminate: true,
		})
	}

	jsonParser := json.NewDecoder(configFile)
	var settings ConfigSettings
	if err = jsonParser.Decode(&settings); err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to parse json file " + ConfigGetFilePath(),
			Severity: "Error",
			Error: err,
			Terminate: true,
		})
	}

	APPLICATION_CONFIG_SETTINGS = settings
}
