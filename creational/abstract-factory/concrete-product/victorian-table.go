package concreteproduct

type VictorianTable struct {
	NoOfLegs int
}

func (v *VictorianTable) GetNumberOfLegs() int {
	return v.NoOfLegs
}

func (v *VictorianTable) GetDescription() string {
	return "This is a victorian table"
}
