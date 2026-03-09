package diff_mode

type DiffMode struct {
	active    bool
	filePath  string
	patchPath string
}

func New(active bool, filePath string, patchPath string) DiffMode {
	return DiffMode{active: active, filePath: filePath, patchPath: patchPath}
}

func (m *DiffMode) Active() bool {
	return m.active
}

func (m *DiffMode) FilePath() string {
	return m.filePath
}

func (m *DiffMode) PatchPath() string {
	return m.patchPath
}

func (m *DiffMode) IsPatchMode() bool {
	return m.patchPath != ""
}

func (m *DiffMode) IsFileMode() bool {
	return m.filePath != ""
}
