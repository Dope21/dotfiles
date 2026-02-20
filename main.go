package main

import (
	"os"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Backup 	bool 		`yaml:"backup"`
	Tools 	[]Tool	`yaml:"tools"`
}

type Tool struct {
	Name				string	`yaml:"name"`
	Description string	`yaml:"description"`
	Conflict		string	`yaml:"conflict"`
	OS					string 	`yaml:"os"`
	LinkMap			LinkMap 	`yaml:"linkmap"`
}

type LinkMap struct {
	Base 			map[string]string `yaml:"-"`
	Windows		map[string]string `yaml:"-"`
	Linux			map[string]string `yaml:"-"`
	MacOS			map[string]string `yaml:"-"`
}

func main() {
	data, err := os.ReadFile("./template.yaml")
	if err != nil {
		log.Fatal("error")
	}

	var config Config
	err = yaml.Unmarshal(data, &config)

	fmt.Printf("YAML data: %+v\n", config)
}