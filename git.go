package main

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func DotFilesDirectoryExists() bool {
	return pathExists(CONFIG_DOTFILES_LOCATION)
}

func GitContext() *git.Repository {
	if !DotFilesDirectoryExists() {
		ConsoleWrite(ConsoleInterface{
			Message: "Cloning repository into " + CONFIG_DOTFILES_LOCATION + " as it could not be found",
			Severity: "Info",
			Error: nil,
			Terminate: false,
		})

		return GitClone()
	}

	if r, err := git.PlainOpen(CONFIG_DOTFILES_LOCATION); err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to open dotfiles repository",
			Severity: "Error",
			Error: err,
			Terminate: true,
		})

		return r
	}

	return nil
}

func GitClone() *git.Repository {
	r, err := git.PlainClone(CONFIG_DOTFILES_LOCATION, false, &git.CloneOptions{
		URL: APPLICATION_CONFIG_SETTINGS.Git.Remote,
		ReferenceName: plumbing.ReferenceName(APPLICATION_CONFIG_SETTINGS.Git.Branch),
	})

	if err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to clone repository",
			Severity: "Error",
			Error: err,
			Terminate: true,
		})
	}

	return r
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

	if !isGitClean(worktree) {
		ConsoleWrite(ConsoleInterface{
			Message: "Dotfiles repository has been left in an unclean stat, ensure working tree is clean before progressing",
			Severity: "Error",
			Error: nil,
			Terminate: true,
		})
	}

	if err := worktree.Pull(&git.PullOptions{}); err != nil {
		ConsoleWrite(ConsoleInterface{
			Message: "Unable to update repository",
			Severity: "Error",
			Error: err,
			Terminate: true,
		})
	}
}