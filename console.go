package main

import (
	"fmt"
	"os"
	"strings"
)

type ConsoleInterface struct {
	Message   string
	Severity  string
	Error     error
	Terminate bool
}

func ConsoleWrite(c ConsoleInterface) {
	if strings.ToLower(c.Severity) == "debug" && !FLAG_VERBOSE {
		return
	}

	if !isEmpty(c.Message) {
		if isEmpty(c.Severity) {
			c.Severity = "Error"
		}

		severity := strings.ToUpper(c.Severity)
		fmt.Printf("==> [%s] %s\n", severity, c.Message)
	}

	if c.Error != nil {
		fmt.Println(c.Error)
	}

	if c.Terminate {
		os.Exit(1)
	}
}
