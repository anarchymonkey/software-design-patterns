package concreteproduct

type VictorianSofa struct {
	NoOfLegs int
}

func (v *VictorianSofa) GetNumberOfLegs() int {
	return v.NoOfLegs
}

func (v *VictorianSofa) GetDescription() string {
	return "This is a victorian sofa"
}
