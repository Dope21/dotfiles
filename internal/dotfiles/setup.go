package dotfiles

import (
	"fmt"
	"path/filepath"
	"slices"

	"github.com/Dope21/dotfiles.git/internal/utils"
)

const BACKUP_PATH = "./backup"

func Setup(configPath string) error {

	logFile, err := utils.CreateLogFile()
	if err != nil {
		fmt.Printf("Warning: Can't create a log file: %v\n", err)
	}

	defer logFile.Close()

	config, err := utils.InitialConfig(configPath)
	if err != nil {
		return err
	}

	runningOS, err := utils.GetPlatform()
	if err != nil {
		return err
	}

	utils.LogAndDisplay("==================================")
	utils.LogAndDisplay("Initial Dotfiles Setup")
	utils.LogAndDisplay("Config file: %s", configPath)
	utils.LogAndDisplay("OS: %s", runningOS)
	utils.LogAndDisplay("==================================")

	for _, tool := range config.Tools {

		utils.LogAndDisplay("")
		utils.LogAndDisplay("==================================")

		if tool.OS != nil && !slices.Contains(tool.OS, runningOS) {
			continue
		}

		utils.LogAndDisplay("")
		utils.LogAndDisplay("Tool: %s", tool.Name)
		utils.LogAndDisplay("Description: %s", tool.Description)
		utils.LogAndDisplay("")

		linkMap := tool.LinkMap.GetOS(runningOS)

		if len(linkMap) > 0 {
			utils.LogAndDisplay("----------------------------------")
			utils.LogAndDisplay("")
			utils.LogAndDisplay("Create Symbolic Link")
			utils.LogAndDisplay("")
		}

		for _, link := range linkMap {
			for source, link := range link {

				utils.LogAndDisplay("Source: %s", source)
				utils.LogAndDisplay("Link: %s", link)

				err := utils.CreateSymlink(source, link, filepath.Join(BACKUP_PATH, tool.Name))
				if err != nil {

					utils.LogAndDisplay("Error: %v", err)

					if tool.Conflict == "skip" {
						utils.LogAndDisplay("Skip mapping")
						continue
					} else {
						return err
					}
				} 
			}
			utils.LogAndDisplay("")
		}

		if len(tool.PostLinkList) > 0 {
			utils.LogAndDisplay("----------------------------------")
			utils.LogAndDisplay("")
			utils.LogAndDisplay("Running Post Symbolic Link Script")
			utils.LogAndDisplay("")
		}

		for _, script := range tool.PostLinkList {
			utils.LogAndDisplay("Script name: %s", script.Name)
			utils.LogAndDisplay("CMD: %s", script.Cmd)

			if err := utils.RunCustomScript(script.Cmd, script.IsPath); err != nil {
				utils.LogAndDisplay("Error: %v", err)
				utils.LogAndDisplay("Skip running this script")
			}
			utils.LogAndDisplay("")
		}

		utils.LogAndDisplay("==================================")
	}

	utils.LogAndDisplay("")
	utils.LogAndDisplay("==================================")
	utils.LogAndDisplay("✅ Setup completed.")
	utils.LogAndDisplay("==================================")

	return nil
}