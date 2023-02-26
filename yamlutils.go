package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigRepositoryObjectToSave struct {
	Path               string `yaml:"path"`
	RelativePathInRepo string `yaml:"relativePathInRepo"`
}
type ConfigRepository struct {
	LocalPath     string                         `yaml:"localPath"`
	ObjectsToSave []ConfigRepositoryObjectToSave `yaml:"objectsToSave"`
}
type Config struct {
	Repositories []ConfigRepository `yaml:"repositories"`
}

func ReadConfig(filename string) *Config {
	configHolder := Config{}

	data, err := os.ReadFile(filename)
	CheckIfError(err)

	err = yaml.Unmarshal(data, &configHolder)
	CheckIfError(err)
	return &configHolder
}
