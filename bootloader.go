package main

import "fmt"

func Bootstrap() {
	LoadApplicationDefaultConfigValues()
	ParseFlags()

	var configFileExists bool = ConfigExists()
	if !configFileExists {
		ConsoleWrite(ConsoleInterface{
			Message: "Config does not exist, please create a " + ConfigGetFilePath(),
			Severity: "Error",
			Terminate: true,
		})
	}

	ParseConfig()
	fmt.Println(APPLICATION_CONFIG_SETTINGS)
}