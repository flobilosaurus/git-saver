package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyObjectsWorksWithDir(t *testing.T) {
	path := CreateTempDirectory()
	subfolderPath := filepath.Join(path, "some", "sub", "folder")
	err := os.MkdirAll(subfolderPath, os.ModePerm)
	check(err)
	repoPath := CreateTempDirectory()
	filename := filepath.Join(subfolderPath, "foo")
	WriteFileWithContent(filename, "bar")

	configRepo := &ConfigRepository{
		LocalPath: repoPath,
		ObjectsToSave: []ConfigRepositoryObjectToSave{
			{
				Path:               filepath.Join(path, "some"),
				RelativePathInRepo: "relative/path/in/repo",
			},
		},
	}

	CopyAllObjectsToRepo(*configRepo)

	filenameInRepo := filepath.Join(repoPath, configRepo.ObjectsToSave[0].RelativePathInRepo, "sub", "folder", "foo")
	exists, err := DoesFileExist(filenameInRepo)
	check(err)

	if !exists {
		t.Fatalf("expected file to exist in repo after copy")
	}
}

func TestCopyObjectsWorksWithFile(t *testing.T) {
	path := CreateTempDirectory()
	repoPath := CreateTempDirectory()
	filename := filepath.Join(path, "foo")
	WriteFileWithContent(filename, "bar")

	configRepo := &ConfigRepository{
		LocalPath: repoPath,
		ObjectsToSave: []ConfigRepositoryObjectToSave{
			{
				Path:               filename,
				RelativePathInRepo: "relative/path/in/repo",
			},
		},
	}

	CopyAllObjectsToRepo(*configRepo)

	filenameInRepo := filepath.Join(repoPath, configRepo.ObjectsToSave[0].RelativePathInRepo, "foo")
	exists, err := DoesFileExist(filenameInRepo)
	check(err)

	if !exists {
		t.Fatalf("expected file to exist in repo after copy")
	}
}
