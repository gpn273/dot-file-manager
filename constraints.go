package main

import (
	"os"
)

const (
	APPLICATION_NAME = "Dot Files Manager"
	APPLICATION_VERSION = "1.0.0"
	APPLICATION_AUTHOR = "Graham Newton"

	CONFIG_DEFAULT_FILE_NAME = "/.dotconfig"
)

var (
	APPLICATION_CONFIG_SETTINGS ConfigSettings
	CONFIG_DEFAULT_FILE_LOCATION string
)

func LoadApplicationDefaultConfigValues() {
	CONFIG_DEFAULT_FILE_LOCATION = os.Getenv("HOME")
}