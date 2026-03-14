package types

import "fmt"

type Config struct {
	Tools []Tool `yaml:"tools,omitempty"`
}

type Tool struct {
	Name           string         `yaml:"name,omitempty"`
	Description    string         `yaml:"description,omitempty"`
	Conflict       string         `yaml:"conflict,omitempty"`
	OS             OSList         `yaml:"os,omitempty"`
	LinkMap        LinkMap        `yaml:"linkmap,omitempty"`
	PostLinkList   []PostLink     `yaml:"post-link,omitempty"`
	MaintenaceList []Maintainance `yaml:"maintenace,omitempty"`
}

type OSList []string

type PostLink struct {
	Name   string   `yaml:"name"`
	IsPath bool     `yaml:"is-path"`
	Cmd    []string `yaml:"cmd"`
}

type Maintainance struct {
	Name   string   `yaml:"name"`
	IsPath bool     `yaml:"is-path"`
	Cmd    []string `yaml:"cmd"`
}

func (c *Config) GetToolByName(name string) (Tool, error) {
	for _, tool := range c.Tools {
		if tool.Name == name {
			return tool, nil
		}
	}

	return Tool{}, fmt.Errorf("can't find tool %q", name)
}

func (t *Tool) GetMaintainScriptByName(name string) (Maintainance, error) {
	for _, script := range t.MaintenaceList {
		if script.Name == name {
			return script, nil
		}
	}

	return Maintainance{}, fmt.Errorf("can't find script %q", name)
}