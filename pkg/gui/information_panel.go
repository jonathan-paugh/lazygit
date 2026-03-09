package gui

import (
	"github.com/jesseduffield/lazygit/pkg/gui/style"
	"github.com/jesseduffield/lazygit/pkg/utils"
)

func (gui *Gui) informationStr() string {
	modePrefix := ""
	if gui.InDiffMode {
		modePrefix = style.FgCyan.Sprintf("-- DIFF MODE -- ")
	} else if gui.InStagingMode {
		modePrefix = style.FgYellow.Sprintf("-- STAGING MODE -- ")
	}

	if activeMode, ok := gui.helpers.Mode.GetActiveMode(); ok {
		return modePrefix + activeMode.InfoLabel()
	}

	return modePrefix + gui.Config.GetVersion()
}

func (gui *Gui) handleInfoClick() error {
	if !gui.g.Mouse {
		return nil
	}

	view := gui.Views.Information

	cx, _ := view.Cursor()
	width := view.Width()

	if activeMode, ok := gui.helpers.Mode.GetActiveMode(); ok {
		if width-cx > utils.StringWidth(gui.c.Tr.ResetInParentheses) {
			return nil
		}
		return activeMode.Reset()
	}

	return nil
}
