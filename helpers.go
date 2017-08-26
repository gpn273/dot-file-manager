package main

import "os"

func pathExists(s string) bool {
	if isEmpty(s) {
		return false
	}

	if _, err := os.Stat(s); os.IsNotExist(err) {
		return false
	}

	return true
}
