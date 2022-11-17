package main

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var testSignature = object.Signature{
	Name:  "John Doe",
	Email: "john@doe.org",
	When:  time.Now(),
}

func TestGetWorktree(t *testing.T) {
	path := CreateTempDirectory()
	repo := InitRepo(path)
	tree := GetWorktree(repo)

	if tree == nil {
		t.Fatalf("worktree was empty")
	}
}

func TestCommitAll(t *testing.T) {
	path := CreateTempDirectory()
	repo := InitRepo(path)
	filename := filepath.Join(path, "foo")
	WriteFileWithContent(filename, "bar")
	CommitAll(repo, "intial commit", &testSignature)
	hasChanges := HasChanges(repo)
	if hasChanges != false {
		t.Fatalf("did not expect changes after commit")
	}

	cIter, err := repo.Log(&git.LogOptions{})
	check(err)

	commitCount := 0
	err = cIter.ForEach(func(c *object.Commit) error {
		commitCount++
		return nil
	})
	check(err)
}

func TestEmptyRepositoryHasNoChanges(t *testing.T) {
	path := CreateTempDirectory()
	repo := InitRepo(path)
	hasChanges := HasChanges(repo)

	if hasChanges != false {
		t.Fatalf("did not expect to have changes")
	}
	DeleteDirectory(path)
}

func TestRepositoryWithModifiedFilesHasChanges(t *testing.T) {
	path := CreateTempDirectory()
	repo := InitRepo(path)
	filename := filepath.Join(path, "foo")
	WriteFileWithContent(filename, "bar")
	CommitAll(repo, "initial commit", &testSignature)
	WriteFileWithContent(filename, "newbar")
	hasChanges := HasChanges(repo)

	if !hasChanges {
		t.Fatalf("did expect to have changes")
	}
	DeleteDirectory(path)
}

func TestRepositoryWithUntrackedFilesHasChanges(t *testing.T) {
	path := CreateTempDirectory()
	repo := InitRepo(path)
	filename := filepath.Join(path, "foo")
	WriteFileWithContent(filename, "bar")
	hasChanges := HasChanges(repo)

	if hasChanges == false {
		t.Fatalf("did expect to have changes")
	}
	DeleteDirectory(path)
}
