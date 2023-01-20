package config

import (
	"errors"

	"github.com/BurntSushi/toml"
)

var Config *ConfigStruct
var Loaded bool = false

type ConfigStruct struct {
	Production bool
	ListenPort int

	Database DbInfo
}

type DbInfo struct {
	Name     string
	Server   string
	Username string
	Password string
	Port     string
}

func NewConfiguration(configFile string) (*ConfigStruct, error) {
	var config = new(ConfigStruct)
	_, err := toml.DecodeFile(configFile, config)

	return config, err
}

func (config *ConfigStruct) Validate() error {
	if config.ListenPort == 0 {
		return errors.New("No ListPort given")
	}

	if len(config.Database.Name) <= 0 {
		return errors.New("No Database.Name given")
	}

	if len(config.Database.Server) <= 0 {
		return errors.New("No Database.Server given")
	}

	if len(config.Database.Username) <= 0 {
		return errors.New("No Database.Username given")
	}

	if len(config.Database.Password) <= 0 {
		return errors.New("No Database.Password given")
	}

	if len(config.Database.Port) <= 0 {
		return errors.New("No Database.Port given")
	}

	return nil
}
