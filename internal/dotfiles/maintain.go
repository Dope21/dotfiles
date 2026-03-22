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

	tool, err := config.GetToolByName(toolName)
	if err != nil {
		return err
	}

	script, err := tool.GetMaintainScriptByName(scriptName) 
	if err != nil {
		return err
	}

	if err := utils.RunCustomScript(script.Cmd, script.IsPath); err != nil {
		return fmt.Errorf("failed to run %s: %w", strings.Join(script.Cmd, " "), err)
	}

	return nil
}