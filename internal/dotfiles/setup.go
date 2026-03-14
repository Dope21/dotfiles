package dotfiles

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/Dope21/dotfiles.git/internal/types"
	"github.com/Dope21/dotfiles.git/internal/utils"
	"gopkg.in/yaml.v3"
)

const BACKUP_PATH = "./backup"

func Setup(configPath string) error {
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

		utils.LogInfo("Start postlink script", true)
		for _, script := range tool.PostLinkList {
			utils.LogInfo(script.Name, false)

			if script.IsPath {
				path, err := filepath.Abs(script.Cmd[0])
				if err != nil {
					return err
				}
				script.Cmd[0] = path
			}

			cmd := exec.Command(script.Cmd[0], script.Cmd[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				return fmt.Errorf("failed to run %s: %w", strings.Join(script.Cmd, " "), err)
			}

		}

		utils.LogInfo("--------------------", true)
	}

	fmt.Println("✅ Setup completed.")
	return nil
}