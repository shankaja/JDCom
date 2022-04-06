package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Swagger struct {
	BaseUrl string `yaml:"base_url"`
	Path    string `yaml:"path"`
}

type Config struct {
	Files   Files `yaml:"files"`
	Swagger Files `yaml:"swagger"`
}

type Files struct {
	JobDescription string `yaml:"job_description"`
	Resume         string `yaml:"resume"`
}

func LoadAll(configPath string) *Config {

	var config Config = Config{}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config
}
