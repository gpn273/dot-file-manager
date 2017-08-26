package main

import (
	"os"
	"path"
)

const (
	APPLICATION_NAME = "Dot Files Manager"
	APPLICATION_VERSION = "1.0.0"
	APPLICATION_AUTHOR = "Graham Newton"
)

var (
	APPLICATION_CONFIG_SETTINGS ConfigSettings
	CONFIG_DEFAULT_FILE_LOCATION string
	CONFIG_DOTFILES_LOCATION string
	CONFIG_DEFAULT_FILE_NAME string
)

func LoadApplicationDefaultConfigValues() {
	CONFIG_DEFAULT_FILE_LOCATION = os.Getenv("HOME")
	CONFIG_DOTFILES_LOCATION = path.Join(CONFIG_DEFAULT_FILE_LOCATION, "/.dotfiles/")
	CONFIG_DEFAULT_FILE_NAME = path.Join(CONFIG_DEFAULT_FILE_LOCATION, "/.dotconfig")
}