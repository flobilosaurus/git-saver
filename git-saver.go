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
	defaultCommitMessage = "updated by https://github.com/flobilosaurus/git-saver"
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
		repoName := GetEqualPadded(configRepo.LocalPath, 40)
		processingRepoOutput := fmt.Sprintf("processing repo '%s' -", repoName)
		fmt.Printf("%s", processingRepoOutput)
		repo := OpenRepo(configRepo.LocalPath)
		Pull(repo)
		CopyAllObjectsToRepo(configRepo)
		if HasChanges(repo) {
			CommitAll(repo, defaultCommitMessage, &object.Signature{
				When: time.Now(),
			})

			fmt.Printf("\r%s commited ", processingRepoOutput)
		} else {
			fmt.Printf("\r%s up to date ", processingRepoOutput)
		}
		wasPushed := Push(repo)
		if wasPushed {
			fmt.Printf("\r%s pushed     \n", processingRepoOutput)
		}
	}
}

func main() {
	app := &cli.App{
		Name:  "git-saver",
		Usage: "automates saving files and folders to git repos",
		Commands: []*cli.Command{
			{
				Name:  "save",
				Usage: "save all configured files to their repositories",
				Action: func(ctx *cli.Context) error {
					SaveAllObjects()
					return nil
				},
			},
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
