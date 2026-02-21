package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Backup 	bool 		`yaml:"backup"`
	Tools 	[]Tool	`yaml:"tools"`
}

type Tool struct {
	Name				string			`yaml:"name"`
	Description string			`yaml:"description"`
	Conflict		string			`yaml:"conflict"`
	OS					[]string 		`yaml:"os"`
	LinkMap			LinkMap 		`yaml:"linkmap"`
}

type LinkMap struct {
	Base		[]map[string]string
	Windows []map[string]string 	`yaml:"windows"`
	Linux 	[]map[string]string 	`yaml:"linux"`
	Macos 	[]map[string]string 	`yaml:"macos"`
}

func (l *LinkMap) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {

	case yaml.SequenceNode:
		var list []map[string]string
		if err := value.Decode(&list); err != nil {
			return err
		}
		l.Base = list
		return nil
	
	case yaml.MappingNode:
		type raw LinkMap
		var grouped raw
		if err := value.Decode(&grouped); err == nil &&
		(grouped.Windows != nil || grouped.Linux != nil || grouped.Macos != nil) {
			l.Windows = grouped.Windows
			l.Linux = grouped.Linux
			l.Macos = grouped.Macos
			return nil
		}
	}
	
	return fmt.Errorf("Invalid type of linkmap")
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