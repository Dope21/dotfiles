package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"gopkg.in/yaml.v3"

	"github.com/Dope21/dotfiles.git/internal/types"
	"github.com/Dope21/dotfiles.git/internal/utils"
)

const BACKUP_PATH = "./backup"

func main() {
	configPath := os.Args[1]
	if err := run(configPath); err != nil {
		log.Fatal(err)
	}
}

func run(configPath string) error {
	utils.LogInfo(fmt.Sprintf("Loading Config from %s", configPath), true)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	// TODO: encforce validation
	var config types.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	runningOS, err := utils.GetPlatform()
	if err != nil {
		return err
	}

	utils.LogInfo(fmt.Sprintf("Running on: %s", runningOS), true)

	utils.LogInfo("--------------------", false)
	for _, tool := range config.Tools {

		if tool.OS != nil && !slices.Contains(tool.OS, runningOS) {
			continue
		}

		utils.LogInfo(tool.Description, true)

		linkMap := tool.LinkMap.GetOS(runningOS)

		utils.LogInfo("Start mapping symlink", true)

		for _, link := range linkMap {
			for source, link := range link {

				utils.LogInfo(fmt.Sprintf("Source: %s", source), false)
				utils.LogInfo(fmt.Sprintf("Link: %s", link), false)

				err := utils.CreateSymlink(source, link)
				if err != nil {

					utils.LogInfo(fmt.Sprintf("Symlink error for %s", tool.Name), false)
					utils.LogInfo(err.Error(), false)

					if tool.Conflict == "skip" {
						utils.LogInfo("Skip mapping", true)
						continue
					} else {
						return err
					}
				} 
			}
		}

		utils.LogInfo("--------------------", true)
	}

	fmt.Println("✅ Setup completed.")
	return nil
}