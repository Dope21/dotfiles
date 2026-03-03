package main

import (
	"log"
	"os"
	"slices"

	"gopkg.in/yaml.v3"

	"github.com/Dope21/dotfiles.git/internal/types"
	"github.com/Dope21/dotfiles.git/internal/utils"
)

const BACKUP_PATH = "./backup"

func main() {

	data, err := os.ReadFile("./template.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config types.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	runningOS, err := utils.GetPlatform()
	if err != nil {
		log.Fatal(err)
	}

	for _, tool := range config.Tools {

		// skip if os to setup is not match
		if tool.OS != nil && !slices.Contains(tool.OS, runningOS) {
			continue
		}

		linkMap := tool.LinkMap.GetOS(runningOS)

		// start mapping
		for _, link := range linkMap {
			for source, link := range link {
				err := utils.CreateSymlink(source, link)
				if err != nil {
					if tool.Conflict == "skip" {
						continue
					} else {
						log.Fatal(err)
					}
				} 
			}
		}
	}
}