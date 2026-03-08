package staging_mode

type StagingMode struct {
	active bool
}

func New(active bool) StagingMode {
	return StagingMode{active: active}
}

func (m *StagingMode) Active() bool {
	return m.active
}
