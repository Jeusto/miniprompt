package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	ShellPrompt     ShellPromptConfig `toml:"shell_prompt"`
	GitBranchConfig struct {
		Disabled bool   `toml:"disabled"`
		Symbol   string `toml:"symbol"`
	} `toml:"git_branch"`
}

type ShellPromptConfig struct {
	PromptSymbol string `toml:"prompt_symbol"`
}

func getConfig() Config {
	var config Config
	if _, err := toml.DecodeFile("miniprompt.toml", &config); err != nil {
		fmt.Println(err)
	}
	return config
}
