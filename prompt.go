package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

const (
	promptSym = "▲"
)

var (
	red    = color("\033[31m%s\033[0m")
	green  = color("\033[32m%s\033[0m")
	cyan   = color("\033[36m%s\033[0m")
	pink   = color("\033[35m%s\033[0m")
	yellow = color("\033[33m%s\033[0m")
)

func color(s string) func(...interface{}) string {
	return func(args ...interface{}) string {
		return fmt.Sprintf(s, fmt.Sprint(args...))
	}
}

func trimPath(cwd, home string) string {
	if strings.HasPrefix(cwd, home) {
		cwd = "~" + strings.TrimPrefix(cwd, home)
	}
	items := strings.Split(cwd, "/")
	for i := 0; i < len(items)-1; i++ {
		items[i] = items[i][:1]
	}
	return filepath.Join(items...)
}

func makePrompt(config Config) string {
	cwd, _ := os.Getwd()
	home := os.Getenv("HOME")
	gitDir := getGitDir()
	promptSym := config.ShellPrompt.PromptSymbol
	gitBranchSym := config.GitBranchConfig.Symbol
	lastExitCode := os.Getenv("EXIT_CODE")
	lastExitCodeSym := ""
	if lastExitCode != "0" {
		lastExitCodeSym = red(fmt.Sprintf("%c %s", '\uf00d', lastExitCode))
	}

	prompt := fmt.Sprintf("\n%s\n%s", cyan(trimPath(cwd, home)), green(promptSym))
	if len(gitDir) > 0 {
		repo, _ := git.PlainOpen(gitDir)
		gitChangedFileCount := fmt.Sprintf("%c %d", '\uf040', gitCountChangedFiles(repo))

		prompt = fmt.Sprintf(
			"\n%s %s %s %s %s\n%s ",
			cyan(trimPath(cwd, home)),
			pink(gitBranchSym),
			pink(gitBranchOrSha(repo)),
			yellow(gitChangedFileCount),
			lastExitCodeSym,
			green(promptSym),
		)
	}
	return prompt
}

func main() {
	config := getConfig()
	fmt.Print(makePrompt(config))
}
