package dotfiles

import (
	"fmt"
	"slices"

	"github.com/Dope21/dotfiles.git/internal/utils"
)

const BACKUP_PATH = "./backup"

func Setup(configPath string) error {

	config, err := utils.InitialConfig(configPath)
	if err != nil {
		return err
	}

	runningOS, err := utils.GetPlatform()
	if err != nil {
		return err
	}

	fmt.Println("==================================")
	fmt.Println("Initial Dotfiles Setup")
	fmt.Printf("Config file: %s\n", configPath)
	fmt.Printf("OS: %s\n", runningOS)
	fmt.Println("==================================")

	for _, tool := range config.Tools {

		fmt.Println()
		fmt.Println("==================================")

		if tool.OS != nil && !slices.Contains(tool.OS, runningOS) {
			continue
		}

		fmt.Println()
		fmt.Printf("Tool: %s\n", tool.Name)
		fmt.Printf("Description: %s\n", tool.Description)
		fmt.Println()
		fmt.Println("----------------------------------")
		fmt.Println()

		linkMap := tool.LinkMap.GetOS(runningOS)

		fmt.Println("Create Symbolic Link")
		fmt.Println()

		for _, link := range linkMap {
			for source, link := range link {

				fmt.Printf("Source: %s\n", source)
				fmt.Printf("Link: %s\n", link)

				err := utils.CreateSymlink(source, link)
				if err != nil {

					fmt.Printf("Error: %s\n", err.Error())

					if tool.Conflict == "skip" {
						fmt.Printf("Skip mapping")
						continue
					} else {
						return err
					}
				} 
			}
			fmt.Println()
		}

		fmt.Println("----------------------------------")
		fmt.Println()

		fmt.Println("Running Post Symbolic Link Script")
		fmt.Println()

		for _, script := range tool.PostLinkList {

			fmt.Printf("Script name: %s\n", script.Name)
			fmt.Printf("CMD: %s\n", script.Cmd)

			if err := utils.RunCustomScript(script.Cmd, script.IsPath); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				fmt.Printf("Skip running this script")
			}
			fmt.Println()
		}

		fmt.Println("==================================")
	}

	fmt.Println()
	fmt.Println("==================================")
	fmt.Println("✅ Setup completed.")
	fmt.Println("==================================")
	return nil
}