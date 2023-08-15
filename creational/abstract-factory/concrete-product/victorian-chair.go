package concreteproduct

type VictorianChair struct {
	NoOfLegs int
}

func (v *VictorianChair) GetNumberOfLegs() int {
	return v.NoOfLegs
}

func (v *VictorianChair) GetDescription() string {
	return "This is a victorian chair"
}
