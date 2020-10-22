package appliances

type Microwave struct {
	typeName string
}

func (sv *Microwave) Start() {
	sv.typeName = " Microwave "
}

func (sv *Microwave) GetPurpose() string {
	return "I am a " + sv.typeName + " I heat food!!"
}
