package concreteproduct

type MordernChair struct {
	NoOfLegs int
}

func (m *MordernChair) GetNumberOfLegs() int {
	return m.NoOfLegs
}

func (m *MordernChair) GetDescription() string {
	return "This is a mordern chair"
}
