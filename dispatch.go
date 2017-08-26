package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

func Dispatch(s string) {
	switch strings.ToLower(s) {
	case "update":
		GitUpdate()
		MapFiles(APPLICATION_CONFIG_SETTINGS.Links, false)
		MapFiles(APPLICATION_CONFIG_SETTINGS.Sources, true)
		break

	default:
		ConsoleWrite(ConsoleInterface{
			Message:   "Unknown action: " + s,
			Severity:  "Error",
			Error:     nil,
			Terminate: true,
		})
	}
}

func MapFiles(files []string, sources bool) {
	if APPLICATION_CONFIG_SETTINGS.Backup {
		CreateBackupDirectory()
	}

	for _, file := range files {
		var originalFile string = path.Join(CONFIG_DEFAULT_FILE_LOCATION, file)
		var vcsFile string = path.Join(CONFIG_DOT_FILES_LOCATION, file)

		if pathExists(originalFile) {
			dateTime := time.Now()
			backupFileName := file + "_" + dateTime.Format("20060102150405")
			CopyFile(originalFile, path.Join(APPLICATION_CONFIG_SETTINGS.BackupPath, backupFileName))
			os.Remove(originalFile)

			ConsoleWrite(ConsoleInterface{
				Message:  fmt.Sprintf("Backing up %s with new name %s to location %s", originalFile, backupFileName, APPLICATION_CONFIG_SETTINGS.BackupPath),
				Severity: "Info",
			})
		}

		os.Symlink(vcsFile, originalFile)
		ConsoleWrite(ConsoleInterface{
			Message:  fmt.Sprintf("Symlinking %s -> %s", vcsFile, originalFile),
			Severity: "Info",
		})

		if sources {
			ConsoleWrite(ConsoleInterface{
				Message: "Run source " + file,
				Severity: "Info",
			})
		}
	}
}

func CreateBackupDirectory() {
	var dir string = APPLICATION_CONFIG_SETTINGS.BackupPath
	if pathExists(dir) {
		return
	}

	os.Mkdir(dir, 0755)
}
