package main

import (
	"os"
)

func Bootstrap() {
	LoadApplicationDefaultConfigValues()
	ParseFlags()

	if flag_show_help {
		ShowHelp()
		os.Exit(1)
	}

	var configFileExists bool = ConfigExists()
	if !configFileExists {
		ConsoleWrite(ConsoleInterface{
			Message: "Config does not exist, please create a " + ConfigGetFilePath(),
			Severity: "Error",
			Terminate: true,
		})
	}

	ParseConfig()
}