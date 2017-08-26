package main

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"os"
)

func DotFilesDirectoryExists() bool {
	return pathExists(CONFIG_DOT_FILES_LOCATION)
}

func GitGetAuth() transport.AuthMethod {
	if !isEmpty(APPLICATION_CONFIG_SETTINGS.Git.Auth.PrivateKey) {
		ConsoleWrite(ConsoleInterface{Message: "Using Private Key authentication", Severity: "Debug"})

		authKey, err := gitssh.NewPublicKeysFromFile("git", APPLICATION_CONFIG_SETTINGS.Git.Auth.PrivateKey, "")
		if err != nil {
			panic(err)
		}

		return authKey
	}

	ConsoleWrite(ConsoleInterface{Message: "Using User and Password authentication", Severity: "Debug"})
	return &gitssh.Password{
		User: APPLICATION_CONFIG_SETTINGS.Git.Auth.UserName,
		Pass: APPLICATION_CONFIG_SETTINGS.Git.Auth.Password,
	}
}

func GitContext() *git.Repository {
	if !DotFilesDirectoryExists() {
		ConsoleWrite(ConsoleInterface{
			Message:   "Cloning repository into " + CONFIG_DOT_FILES_LOCATION + " as it could not be found",
			Severity:  "Info",
			Error:     nil,
			Terminate: false,
		})

		GitClone()
	}

	r, err := git.PlainOpen(CONFIG_DOT_FILES_LOCATION)
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message:   "Unable to open dotfiles repository",
			Severity:  "Error",
			Error:     err,
			Terminate: true,
		})
	}

	r.Fetch(&git.FetchOptions{
		RemoteName: APPLICATION_CONFIG_SETTINGS.Git.RemoteName,
		Auth: GitGetAuth(),
		Progress: os.Stdout,
	})

	return r
}

func GitClone() {
	_, err := git.PlainClone(CONFIG_DOT_FILES_LOCATION, false, &git.CloneOptions{
		URL:        APPLICATION_CONFIG_SETTINGS.Git.Remote,
		RemoteName: APPLICATION_CONFIG_SETTINGS.Git.RemoteName,
		Auth:       GitGetAuth(),
		Progress:   os.Stdout,
	})

	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message:   "Unable to clone repository",
			Severity:  "Error",
			Error:     err,
			Terminate: true,
		})
	}
}

func GitWorkingTree(r *git.Repository) *git.Worktree {
	workTree, err := r.Worktree()
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message:   "Unable to open repository working tree",
			Severity:  "Error",
			Error:     err,
			Terminate: true,
		})
	}

	return workTree
}

func isGitClean(r *git.Worktree) bool {
	status, err := r.Status()
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message:   "Unable to determine repository status",
			Severity:  "Error",
			Error:     err,
			Terminate: true,
		})
	}

	return status.IsClean()
}

func GitUpdate() {
	ctx := GitContext()
	worktree := GitWorkingTree(ctx)

	if !isGitClean(worktree) {
		ConsoleWrite(ConsoleInterface{
			Message:   "Dotfiles repository has been left in an unclean stat, ensure working tree is clean before progressing",
			Severity:  "Error",
			Error:     nil,
			Terminate: true,
		})
	}

	remotes, err := ctx.Remotes()
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message:   "Unable to determine the repositories remote",
			Severity:  "Error",
			Error:     err,
			Terminate: true,
		})
	}

	worktree.Pull(&git.PullOptions{
		RemoteName: remotes[0].Config().Name,
		Auth:       GitGetAuth(),
		Progress:   os.Stdout,
	})
}
