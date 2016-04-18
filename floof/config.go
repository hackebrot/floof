package main

import (
	"os/user"
	"path"

	"github.com/BurntSushi/toml"
)

//UserConfig with Gist information
type UserConfig struct {
	Gist gist
}

//gist representing the according section in the user config
type gist struct {
	Username string
	Token    string
}

//getFile returns the current user's config file path
func getFile() (h string, err error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tomlFile := path.Join(usr.HomeDir, ".floofrc")
	return tomlFile, nil
}

//LoadConfig the given toml file to a UserConfig
func LoadConfig() (*UserConfig, error) {
	tomlFile, err := getFile()
	if err != nil {
		return nil, err
	}

	var c UserConfig
	if _, err := toml.DecodeFile(tomlFile, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
