package memento

type Memento struct {
	state string
}

func (m *Memento) GetSavedState() string {
	return m.state
}
