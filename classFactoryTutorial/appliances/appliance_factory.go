package appliances

//Factory Patter: All the Objects are created from here
import "errors"

type Appliance interface {
	Start()
	GetPurpose() string
}

// IOTA, means all the const variables are automatically incremented e.g.: STOVE = 0, FRIDGE = 1
const (
	STOVE = iota
	FRIDGE
	MICROWAVE
)

func CreateAppliance(myType int) (Appliance, error) {
	switch myType {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	case MICROWAVE:
		return new(Microwave), nil
	default:
		return nil, errors.New("Invalid Appliance Type")
	}
}
