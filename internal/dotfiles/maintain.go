package dotfiles

import (
	"fmt"
	"strings"

	"github.com/Dope21/dotfiles.git/internal/utils"
)

func Maintain(toolName, scriptName string) error {
	config, err := utils.GetConfig()
	if err != nil {
		return err
	}

	for _, tool := range config.Tools {

		if tool.Name != toolName {
			continue
		}

		for _, script := range tool.MaintenaceList {

			if script.Name != scriptName {
				continue
			}

			if err := utils.RunCustomScript(script.Cmd, script.IsPath); err != nil {
				return fmt.Errorf("failed to run %s: %w", strings.Join(script.Cmd, " "), err)
			}
		}
	}

	return nil
}