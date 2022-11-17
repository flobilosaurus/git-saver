package main

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateTempDirectory() string {
	d, err := os.MkdirTemp("", "test_repo")
	check(err)
	return d
}

func InitRepo(path string) *git.Repository {
	r, err := git.PlainInit(path, false)
	check(err)
	return r
}

func DeleteDirectory(path string) {
	defer os.RemoveAll(path)
}

func WriteFileWithContent(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0755)
	check(err)
}
