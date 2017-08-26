package main

import (
	"gopkg.in/src-d/go-git.v4"
	"fmt"
	"golang.org/x/crypto/ssh"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"io/ioutil"
)

func DotFilesDirectoryExists() bool {
	return pathExists(CONFIG_DOTFILES_LOCATION)
}

func GitGetAuth() *gitssh.PublicKeys {
	privateKey, err := ioutil.ReadFile(APPLICATION_CONFIG_SETTINGS.Git.PrivateKey)
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to open private key file",
			Severity: "Error",
			Error: err,
			Terminate: true,
		})
	}

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to parse private key",
			Severity: "Error",
			Error: err,
			Terminate: true,
		})
	}

	return &gitssh.PublicKeys{User: APPLICATION_CONFIG_SETTINGS.Git.User, Signer: signer}
}

func GitContext() *git.Repository {
	if !DotFilesDirectoryExists() {
		ConsoleWrite(ConsoleInterface{
			Message: "Cloning repository into " + CONFIG_DOTFILES_LOCATION + " as it could not be found",
			Severity: "Info",
			Error: nil,
			Terminate: false,
		})

		GitClone()
	}

	r, err := git.PlainOpen(CONFIG_DOTFILES_LOCATION)
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to open dotfiles repository",
			Severity: "Error",
			Error: err,
			Terminate: true,
		})
	}

	return r
}

func GitClone() {
	_, err := git.PlainClone(CONFIG_DOTFILES_LOCATION, false, &git.CloneOptions{
		URL: APPLICATION_CONFIG_SETTINGS.Git.Remote,
		Auth: GitGetAuth(),
	})

	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to clone repository",
			Severity: "Error",
			Error: err,
			Terminate: true,
		})
	}
}

func GitWorkingTree(r *git.Repository) *git.Worktree {
	worktree, err := r.Worktree()
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to open repository working tree",
			Severity: "Error",
			Error: err,
			Terminate: true,
		})
	}

	return worktree
}

func isGitClean(r *git.Worktree) bool {
	status, err := r.Status()
	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to determine repository status",
			Severity: "Error",
			Error: err,
			Terminate: true,
		})
	}

	return status.IsClean()
}

func GitUpdate() {
	ctx := GitContext()
	worktree := GitWorkingTree(ctx)
	fmt.Println(worktree)

	if !isGitClean(worktree) {
		ConsoleWrite(ConsoleInterface{
			Message: "Dotfiles repository has been left in an unclean stat, ensure working tree is clean before progressing",
			Severity: "Error",
			Error: nil,
			Terminate: true,
		})
	}

	if err := worktree.Pull(&git.PullOptions{
		Auth: GitGetAuth(),
	}); err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to update repository",
			Severity: "Error",
			Error: err,
			Terminate: true,
		})
	}
}