package main

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/kevinburke/ssh_config"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("error aaa: %s", err)
}

func GetWorktree(repo *git.Repository) *git.Worktree {
	tree, err := repo.Worktree()
	CheckIfError(err)

	return tree
}

func HasChanges(repo *git.Repository) bool {
	tree := GetWorktree(repo)
	_, err := tree.Add(".")
	CheckIfError(err)
	status, err := tree.Status()
	CheckIfError(err)
	count := 0
	for range status {
		count++
	}

	return count > 0
}

func OpenRepo(path string) *git.Repository {
	expandedPath := expand(path)
	repo, err := git.PlainOpen(expandedPath)
	CheckIfError(err)
	return repo
}

func CommitAll(repo *git.Repository, message string, signature *object.Signature) {
	tree := GetWorktree(repo)
	_, err := tree.Add(".")
	CheckIfError(err)
	commit, err := tree.Commit(message, &git.CommitOptions{
		All: true,
		Author: &object.Signature{
			Name: "git-saver",
			When: time.Now(),
		},
	})
	CheckIfError(err)
	_, err = repo.CommitObject(commit)
	CheckIfError(err)
}

func Pull(repo *git.Repository) {
	tree := GetWorktree(repo)
	publicKeys := GetAuth(repo)
	err := tree.Pull(&git.PullOptions{RemoteName: "origin", Auth: publicKeys})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		fmt.Printf("error while pulling: %s", err)
	}
}

func Push(repo *git.Repository) (hasPushed bool) {
	publicKeys := GetAuth(repo)
	err := repo.Push(&git.PushOptions{
		Auth: publicKeys,
	})
	if err != nil {
		if !errors.Is(err, git.NoErrAlreadyUpToDate) {
			fmt.Printf("error000: %s", err)
		}
		return false
	}

	return true
}

func GetAuth(repo *git.Repository) *ssh.PublicKeys {
	host := GetHostOfOriginRemote(repo)
	privateKeyFilepath := ssh_config.Get(host, "IdentityFile")
	user := ssh_config.Get(host, "User")
	expandedPath := expand(privateKeyFilepath)
	publicKeys, err := ssh.NewPublicKeysFromFile(user, expandedPath, "")
	CheckIfError(err)
	return publicKeys
}

func GetHostOfOriginRemote(repo *git.Repository) string {
	cfg, err := repo.Config()
	CheckIfError(err)
	r, _ := regexp.Compile("@(.*)+:")
	host := r.FindStringSubmatch(cfg.Remotes["origin"].URLs[0])
	return host[1]
}
