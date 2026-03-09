package types

import (
	"github.com/jesseduffield/lazygit/pkg/gui/modes/cherrypicking"
	"github.com/jesseduffield/lazygit/pkg/gui/modes/diffing"
	"github.com/jesseduffield/lazygit/pkg/gui/modes/filtering"
	"github.com/jesseduffield/lazygit/pkg/gui/modes/diff_mode"
	"github.com/jesseduffield/lazygit/pkg/gui/modes/marked_base_commit"
	"github.com/jesseduffield/lazygit/pkg/gui/modes/staging_mode"
)

type Modes struct {
	Filtering        filtering.Filtering
	CherryPicking    *cherrypicking.CherryPicking
	Diffing          diffing.Diffing
	MarkedBaseCommit marked_base_commit.MarkedBaseCommit
	StagingMode      staging_mode.StagingMode
	DiffMode         diff_mode.DiffMode
}
