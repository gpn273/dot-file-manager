package main

import (
	"os"
	"encoding/json"
	"path"
)

type GitAuthSettings struct {
	PrivateKey string `json:"private_key"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type GitSettings struct {
	Remote string `json:"remote"`
	RemoteName string `json:"remote_name"`
	Branch string `json:"branch"`
	Auth GitAuthSettings `json:"auth"`
}

type ConfigSettings struct {
	Git GitSettings `json:"git"`
	Backup bool `json:"backup"`
}

func ConfigGetFilePath() string  {
	return path.Join(CONFIG_DEFAULT_FILE_LOCATION, CONFIG_DEFAULT_FILE_NAME)
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
