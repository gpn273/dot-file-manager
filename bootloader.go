package main

import (
	"os"
)

func Bootstrap() {
	LoadApplicationDefaultConfigValues()
	ParseFlags()
	ParseArgs()

	if FLAG_SHOW_HELP {
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