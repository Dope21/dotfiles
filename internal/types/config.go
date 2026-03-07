package types

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Tools  []Tool `yaml:"tools,omitempty"`
}

type Tool struct {
	Name        	string  		`yaml:"name,omitempty"`
	Description 	string  		`yaml:"description,omitempty"`
	Conflict    	string  		`yaml:"conflict,omitempty"`
	OS          	OSList  		`yaml:"os,omitempty"`
	LinkMap     	LinkMap 		`yaml:"linkmap,omitempty"`
	PostLinkList 	[]PostLink 	`yaml:"post-link,omitempty"`
}

type LinkMap struct {
	Base    []map[string]string
	Windows []map[string]string `yaml:"windows,omitempty"`
	Linux   []map[string]string `yaml:"linux,omitempty"`
	Macos   []map[string]string `yaml:"macos,omitempty"`
}

type OSList []string

type PostLink struct {
	Name  	string 		`yaml:"name"`
	IsPath 	bool			`yaml:"is-path"`
	Cmd 		[]string 	`yaml:"cmd"`
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

func (l *LinkMap) GetOS(os string) []map[string]string {
	if len(l.Base) != 0 {
		return l.Base
	}

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