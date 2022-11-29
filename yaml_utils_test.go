package main

import (
	"path/filepath"
	"testing"
)

func AssertEqual(t *testing.T, first string, second string) {
	if first != second {
		t.Fatalf("expected %s but was %s", first, second)
	}
}

func TestReadConfig(t *testing.T) {
	path := CreateTempDirectory()
	configFilename := filepath.Join(path, ".gitsaver-config.yaml")
	WriteFileWithContent(configFilename, defaultConfigContent)
	config := ReadConfig(configFilename)

	if len(config.Repositories) != 1 {
		t.Fatalf("expected 1 repositories but found %d", len(config.Repositories))
	}

	AssertEqual(t, config.Repositories[0].LocalPath, "~/repositories/dotfiles")
	AssertEqual(t, config.Repositories[0].ObjectsToSave[0].Path, "~/.config/nvim")
	AssertEqual(t, config.Repositories[0].ObjectsToSave[0].RelativePathInRepo, "my-nvim-conf")
}
