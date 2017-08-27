package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

func Dispatch(s string) {
	switch GetFriendlyArgName() {
	case "update":
		GitUpdate()
		MapFiles(APPLICATION_CONFIG_SETTINGS.Links, false)
		MapFiles(APPLICATION_CONFIG_SETTINGS.Sources, true)
		break

	case "create-config":
		CreateConfig()
		break

	case "purge-backups":
		if err := PurgeBackups(); err != nil {
			ConsoleWrite(ConsoleInterface{
				Error:     err,
				Terminate: true,
			})
		}
		break

	case "create-global-installation":
		CreateGlobalInstallation()
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

	for _, f := range files {
		var originalFile string = path.Join(CONFIG_DEFAULT_FILE_LOCATION, f)
		var vcsFile string = path.Join(CONFIG_DOT_FILES_LOCATION, f)

		if pathExists(originalFile) {
			dateTime := time.Now()
			backupFileName := f + "_" + dateTime.Format("20060102150405")
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
				Message:  "Run source " + f,
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

func CreateConfig() {
	if exists := ConfigExists(); exists {
		ConsoleWrite(ConsoleInterface{
			Message:   "There is already config located: " + CONFIG_DEFAULT_FILE_NAME,
			Severity:  "Warn",
			Terminate: true,
		})
	}

	var conf ConfigSettings = ConfigSettings{
		Git: GitSettings{
			Remote:     "git url",
			RemoteName: "origin",
			Branch:     "master",
			Auth: GitAuthSettings{
				PrivateKey: path.Join(CONFIG_DEFAULT_FILE_LOCATION, "/.ssh/"),
				UserName:   "git username",
				Password:   "git password",
			},
		},
		Backup:     true,
		BackupPath: path.Join(CONFIG_DEFAULT_FILE_LOCATION, "/.dotfiles_backup/"),
		Links: []string{
			".bashrc",
		},
		Sources: []string{
			".macos",
		},
	}

	jsonPayload, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Error:     err,
			Terminate: true,
		})
	}

	f, err := os.Create(CONFIG_DEFAULT_FILE_NAME)
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Error:     err,
			Terminate: true,
		})
	}

	defer f.Close()

	if _, err := f.Write(jsonPayload); err != nil {
		ConsoleWrite(ConsoleInterface{
			Error:     err,
			Terminate: true,
		})
	}

	if err := f.Sync(); err != nil {
		ConsoleWrite(ConsoleInterface{
			Error:     err,
			Terminate: true,
		})
	}

	ConsoleWrite(ConsoleInterface{
		Message:   "Successfully created example config file located: " + CONFIG_DEFAULT_FILE_NAME,
		Severity:  "Info",
		Terminate: true,
	})
}

func PurgeBackups() error {
	d, err := os.Open(APPLICATION_CONFIG_SETTINGS.BackupPath)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(APPLICATION_CONFIG_SETTINGS.BackupPath, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateGlobalInstallation() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Error:     err,
			Terminate: true,
		})
	}

	var globalPath string = path.Join("/usr/local/bin/", os.Args[0])
	var selfPath string = path.Join(dir, os.Args[0])

	if err := os.Symlink(selfPath, globalPath); err != nil {
		ConsoleWrite(ConsoleInterface{
			Message:   "Failed to create global installation",
			Error:     err,
			Terminate: true,
		})
	}

	ConsoleWrite(ConsoleInterface{
		Message:   "Successfully create global installation in " + globalPath,
		Severity:  "Success",
		Terminate: true,
	})
}
