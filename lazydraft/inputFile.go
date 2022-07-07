package lazydraft

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

const projectFileDoesNotExistsError = "\nprojects.yml file could not be found. See 'lazydraft help init'"
const projectFormatInvalidError = "\nprojects.yml file format is invalid. See 'lazydraft config'"
const ConfigFileName = ".config/lazydraft/projects.yml"
const ConfigFileDir = "lazydraft"
const ConfigBaseDir = ".config"
const userHomeDirectoryError = "user home directory cannot be retrieved"
const configFilePathError = "lazydraft file path cannot be retrieved"

type InputFile struct {
	Path string
}

type ProjectConfig struct {
	Data map[string]YamlFile
}

func (yf InputFile) readProjectConfig() (*ProjectConfig, error) {
	data, err := ioutil.ReadFile(yf.Path)
	if err != nil {
		return nil, errors.New(projectFileDoesNotExistsError)
	}
	projectConfig := ProjectConfig{
		Data: make(map[string]YamlFile),
	}
	err = yaml.Unmarshal(data, &projectConfig.Data)
	if err != nil {
		return nil, errors.New(projectFormatInvalidError)
	}
	return &projectConfig, nil
}

func GetProjectConfig() (*ProjectConfig, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return nil, errors.New(configFilePathError)
	}
	projectConfig, err := configPath.readProjectConfig()
	if err != nil {
		return nil, err
	}
	return projectConfig, nil
}

func getConfigFilePath() (*InputFile, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.New(userHomeDirectoryError)
	}
	configFile := InputFile{
		Path: dirname + "/" + ConfigFileName,
	}
	return &configFile, nil
}
