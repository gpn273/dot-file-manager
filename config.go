package main

import (
	"encoding/json"
	"os"
)

type GitAuthSettings struct {
	PrivateKey string `json:"private_key"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
}

type GitSettings struct {
	Remote     string `json:"remote"`
	RemoteName string `json:"remote_name"`
	Branch     string `json:"branch"`
	Auth       GitAuthSettings `json:"auth"`
}

type ConfigSettings struct {
	Git        GitSettings `json:"git"`
	Backup     bool `json:"backup"`
	BackupPath string `json:"backup_path"`
	Links      []string `json:"links"`
	Sources []string `json:"sources"`
}

func ConfigExists() bool {
	return pathExists(CONFIG_DEFAULT_FILE_NAME)
}

func ParseConfig() {
	configFile, err := os.Open(CONFIG_DEFAULT_FILE_NAME)
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message:   "Unable to open config file: " + err.Error(),
			Severity:  "Error",
			Terminate: true,
		})
	}

	jsonParser := json.NewDecoder(configFile)
	var settings ConfigSettings
	if err = jsonParser.Decode(&settings); err != nil {
		ConsoleWrite(ConsoleInterface{
			Message:   "Unable to parse json file " + CONFIG_DEFAULT_FILE_NAME,
			Severity:  "Error",
			Error:     err,
			Terminate: true,
		})
	}

	APPLICATION_CONFIG_SETTINGS = settings
}
