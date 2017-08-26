package main

import "strings"

func Dispatch(s string) {
	switch strings.ToLower(s) {
	case "update":
		GitUpdate()
		break

	default:
		ConsoleWrite(ConsoleInterface{
			Message: "Unknown action: " + s,
			Severity: "Error",
			Error: nil,
			Terminate: true,
		})
	}
}
