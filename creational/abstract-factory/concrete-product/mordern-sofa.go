package concreteproduct

type MordernSofa struct {
	NoOfLegs int
}

func (m *MordernSofa) GetNumberOfLegs() int {
	return m.NoOfLegs
}

func (m *MordernSofa) GetDescription() string {
	return "This is a mordern chair"
}
