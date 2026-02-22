package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"slices"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Backup 	bool 		`yaml:"backup,omitempty"`
	Tools 	[]Tool	`yaml:"tools,omitempty"`
}

type Tool struct {
	Name				string			`yaml:"name,omitempty"`
	Description string			`yaml:"description,omitempty"`
	Conflict		string			`yaml:"conflict,omitempty"`
	OS					OSList 			`yaml:"os,omitempty"`
	LinkMap			LinkMap 		`yaml:"linkmap,omitempty"`
}

type LinkMap struct {
	Base		[]map[string]string
	Windows []map[string]string 	`yaml:"windows,omitempty"`
	Linux 	[]map[string]string 	`yaml:"linux,omitempty"`
	Macos 	[]map[string]string 	`yaml:"macos,omitempty"`
}

type OSList []string

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

func (l *LinkMap) GetOS(os string) []map[string]string {
	switch os {
	case "windows":
		return l.Windows
	case "linux":
		return l.Linux
	case "darwin":
		return l.Macos
	default:
		return l.Base
	}
}

const BACKUP_PATH = "./backup"
const LOGS_PATH = "./logs"

func getPlatform() (string, error) {
	os := runtime.GOOS

	switch os {
	case "windows", "linux", "darwin":
		return os, nil
	default:
		return "", fmt.Errorf("%s OS is not supported", runtime.GOOS)
	}
}

func main() {


	data, err := os.ReadFile("./template.yaml")
	if err != nil {
		log.Fatal("error")
	}

	var config Config
	err = yaml.Unmarshal(data, &config)

	runningOS, err := getPlatform()
	if err != nil {
		log.Fatal("%w", err)
	}

	fmt.Println("running on", runningOS)

	for index, tool := range config.Tools {

		fmt.Printf("Index %d, Tool Name: %s\n", index, tool.Name)

		// check os config
		fmt.Println("os to setup", tool.OS)

		// skip if os to setup is not match
		if tool.OS != nil && !slices.Contains(tool.OS, runningOS) {
			continue
		}

		linkMap := tool.LinkMap.GetOS(runningOS)

		// start mapping
		for _, link := range linkMap {
			for src, dst := range link {
				fmt.Println(src, dst)
			}
		}
	}
}