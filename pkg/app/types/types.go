package app

import (
	integrationTypes "github.com/jesseduffield/lazygit/pkg/integration/types"
)

// StartArgs is the struct that represents some things we want to do on program start
type StartArgs struct {
	// GitArg determines what context we open in
	GitArg GitArg
	// integration test (only relevant when invoking lazygit in the context of an integration test)
	IntegrationTest integrationTypes.IntegrationTest
	// FilterPath determines which path we're going to filter on so that we only see commits from that file.
	FilterPath string
	// ScreenMode determines the initial Screen Mode (normal, half or full) to use
	ScreenMode string
	// Mode determines the operating mode (e.g. "staging" for restricted staging-only mode)
	Mode string
	// DiffFile is the path to a specific file to diff (only used with --mode=diff)
	DiffFile string
	// DiffPatch is the path to a patch file to display (only used with --mode=diff)
	DiffPatch string
}

type GitArg string

const (
	GitArgNone   GitArg = ""
	GitArgStatus GitArg = "status"
	GitArgBranch GitArg = "branch"
	GitArgLog    GitArg = "log"
	GitArgStash  GitArg = "stash"
)

const ModeStagingValue = "staging"
const ModeDiffValue = "diff"

func NewStartArgs(filterPath string, gitArg GitArg, screenMode string, mode string, diffFile string, diffPatch string, test integrationTypes.IntegrationTest) StartArgs {
	return StartArgs{
		FilterPath:      filterPath,
		GitArg:          gitArg,
		ScreenMode:      screenMode,
		Mode:            mode,
		DiffFile:        diffFile,
		DiffPatch:       diffPatch,
		IntegrationTest: test,
	}
}
