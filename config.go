package main

import (
	"fmt"
	"io/ioutil"
	"errors"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Apps []string `yaml:"apps"`
	ShowNotifications bool `yaml:"show_notifications"`
}

func loadConfig(fpath string) *Config {
	buf, err := ioutil.ReadFile(fpath)

	if err != nil {
		exit(errors.New(fmt.Sprintf("Failed to read config file from %v", fpath)))
		return nil;
	}

	cfg := &Config{}

	err = yaml.Unmarshal(buf, cfg)

	if err != nil {
		exit(errors.New(fmt.Sprintf("Failed to load parse YAML from %v", fpath)))
		return nil;
	}
	return cfg;
}
