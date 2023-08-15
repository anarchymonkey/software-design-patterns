package concreteproduct

type MordernTable struct {
	NoOfLegs int
}

func (m *MordernTable) GetNumberOfLegs() int {
	return m.NoOfLegs
}

func (m *MordernTable) GetDescription() string {
	return "This is a mordern chair"
}
