package main

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func getGitDir() string {
	cwd, _ := os.Getwd()
	for {
		dirs, _ := os.ReadDir(cwd)
		for _, d := range dirs {
			if ".git" == d.Name() {
				return cwd
			} else if cwd == "/" {
				return ""
			}
		}
		cwd = filepath.Dir(cwd)
	}
}

// Returns the current git branch or current ref sha.
func gitBranchOrSha(repo *git.Repository) string {
	h, _ := repo.Head()
	head := h.Strings()[0]
	sha := h.Strings()[1]
	if branch := filepath.Base(head); branch != "HEAD" {
		return branch
	} else {
		// Detached HEAD state; return the first 7
		// chars of commit sha.
		return sha[:7]
	}
}

func gitStatus(repo *git.Repository) string {
	worktree, _ := repo.Worktree()
	status, _ := worktree.Status()
	if status.IsClean() {
		return green("•")
	}
	return red("×")
}

func gitCountChangedFiles(repo *git.Repository) int {
	worktree, _ := repo.Worktree()
	status, _ := worktree.Status()

	count := 0
	for _, entry := range status {
		if entry.Staging != git.Untracked {
			count++
		}
	}

	return count
}
