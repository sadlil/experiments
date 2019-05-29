package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type AppInfo struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Config struct {
	App AppInfo `yaml:"app"`
}

type TestUrl struct {
	Name      string `yaml:"name"`
	Url       string `yaml:"url"`
	ProxyPort string `yaml:"proxy_port,omitempty"`
}

type TestUrls struct {
	ATest []TestUrl `yaml:"test_url"`
}

func main() {
	filename, _ := filepath.Abs("files/1.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	var test TestUrls
	err = yaml.Unmarshal(yamlFile, &test)
	if err != nil {
		panic(err)
	}

	fmt.Println("Application : ", config.App.Name, "\nVersion : ", config.App.Version)
	fmt.Println(test)
}
