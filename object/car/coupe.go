package car

type Coupe struct {
	Car
}

func (c Coupe) Speed() int {
	return 100
}
