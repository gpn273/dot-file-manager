package main

import (
	"fmt"
	"strings"
	"os"
)

type ConsoleInterface struct {
	Message   string
	Severity  string
	Error error
	Terminate bool
}

func ConsoleWrite(c ConsoleInterface) {
	severity := strings.ToUpper(c.Severity)
	fmt.Printf("==> [%s] %s\n", severity, c.Message)

	if c.Error != nil {
		fmt.Println(c.Error)
	}

	if c.Terminate {
		os.Exit(1)
	}
}