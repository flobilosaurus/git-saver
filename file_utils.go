package main

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	cp "github.com/otiai10/copy"
)

// to expand tilde in directories
func expand(inputPath string) string {
	usr, err := user.Current()
	CheckIfError(err)
	dir := usr.HomeDir
	if inputPath == "~" {
		// In case of "~", which won't be caught by the "else if"
		return dir
	} else if strings.HasPrefix(inputPath, "~/") {
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		return filepath.Join(dir, inputPath[2:])
	}
	return inputPath
}

func DoesFileExist(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return true, err
	}
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

func CopyAllObjectsToRepo(configRepo ConfigRepository) {
	repoPath := expand(configRepo.LocalPath)
	for _, objectToSave := range configRepo.ObjectsToSave {
		absolutePathObjectToSave := expand(objectToSave.Path)
		isDir, err := isDirectory(absolutePathObjectToSave)
		CheckIfError(err)
		targetPath := filepath.Join(repoPath, objectToSave.RelativePathInRepo)
		if !isDir {
			targetPath = filepath.Join(targetPath, filepath.Base(absolutePathObjectToSave))
		}
		err = cp.Copy(absolutePathObjectToSave, targetPath)
		CheckIfError(err)
	}
}
