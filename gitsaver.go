package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/urfave/cli/v2"
)

var (
	configFileName       = ".git-saver.config.yaml"
	defaultCommitMessage = "updated by git-saver"
	defaultConfigContent = `
repositories: 
  - localPath: ~/repositories/dotfiles 
    objectsToSave:
    - path: ~/.config/nvim
      relativePathInRepo: my-nvim-conf 
`
)

func GetConfigPath() string {
	userHomePath, err := os.UserHomeDir()
	CheckIfError(err)
	return filepath.Join(userHomePath, configFileName)
}

func CreateInitialConfig(configFilePath string) {
	configExists, err := DoesFileExist(configFilePath)
	CheckIfError(err)

	if !configExists {
		err = os.WriteFile(configFilePath, []byte(defaultConfigContent), 0755)
		CheckIfError(err)
	} else {
		log.Printf("Did not create config as it already exists under: %s", configFilePath)
	}
}

func SaveAllObjects() {
	configPath := GetConfigPath()
	config := ReadConfig(configPath)
	for _, configRepo := range config.Repositories {
		fmt.Printf("processing repo: %s\n", configRepo.LocalPath)
		repo := OpenRepo(configRepo.LocalPath)
		CopyAllObjectsToRepo(configRepo)
		if HasChanges(repo) {
			CommitAll(repo, defaultCommitMessage, &object.Signature{
				When: time.Now(),
			})

			fmt.Printf("- successfully updated %s\n", configRepo.LocalPath)
		} else {
			fmt.Printf("- already up to date\n")
		}
		fmt.Printf("- pushing commits\n")
		Push(repo)
	}
}

func main() {
	app := &cli.App{
		Name:  "git-saver",
		Usage: "single call to commit and push files to configured",
		Action: func(ctx *cli.Context) error {
			SaveAllObjects()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "init-config",
				Usage: "create initial config in users home directory",
				Action: func(cCtx *cli.Context) error {
					CreateInitialConfig(GetConfigPath())
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
