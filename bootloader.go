package main

import (
	"os"
)

func Bootstrap() {
	LoadApplicationDefaultConfigValues()
	ParseFlags()
	ParseArgs()

	if *FLAG_SHOW_HELP {
		ShowHelp()
		os.Exit(1)
	}

	if GetFriendlyArgName() != "create-config" {
		var configFileExists bool = ConfigExists()
		if !configFileExists {
			ConsoleWrite(ConsoleInterface{
				Message: "Config does not exist, please create a " + CONFIG_DEFAULT_FILE_NAME,
				Severity: "Error",
				Terminate: true,
			})
		}

		ParseConfig()
	}
}