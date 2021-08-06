package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	SrcPath string `yaml:"srcPath"`
}

var ConfigEntry Config

func LoadConfig(path string) (setting Config) {
	if path == "" {

	}
	config, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &setting)
	ConfigEntry = setting
	return
}
