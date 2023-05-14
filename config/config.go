package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Server struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Protocol string `yaml:"protocol"`
	} `yaml:"server"`
	Jwt struct {
		Key string `yaml:"key"`
	} `yaml:"jwt"`
	Hash struct {
		Secret string `yaml:"secret"`
	}
	Postgres struct {
		Address string `yaml:"address"`
	} `yaml:"postgres"`
	Source struct {
		Name string `yaml:"name"`
	} `yaml:"source"`
}

var Configure *Config

func GetConfig(fileName string) (*Config, error) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", "config.yaml", err)
	}

	return c, nil
}
