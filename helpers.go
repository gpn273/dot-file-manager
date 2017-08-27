package main

import (
	"io"
	"os"
)

func pathExists(s string) bool {
	if isEmpty(s) {
		return false
	}

	if _, err := os.Stat(s); os.IsNotExist(err) {
		return false
	}

	return true
}

func isEmpty(s string) bool {
	return len(s) == 0
}

func CopyFile(originalFile, destFile string) {
	srcFile, err := os.Open(originalFile)
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message:   "Unable to copy file",
			Error:     err,
			Terminate: true,
		})
	}
	defer srcFile.Close()

	destCtxFile, err := os.Create(destFile) // creates if file doesn't exist
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message:   "Unable to create destination file",
			Error:     err,
			Terminate: true,
		})
	}
	defer destCtxFile.Close()

	_, err = io.Copy(destCtxFile, srcFile) // check first var for number of bytes copied
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message:   "Failed to copy file contents to new location",
			Error:     err,
			Terminate: true,
		})
	}
}
