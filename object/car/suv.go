package car


import "math/rand"

/**
 * Anonymous field example
 */
type Suv struct {
	Car
}

func (s Suv) Speed() int {
	r := rand.New(rand.NewSource(99))
	return r.Int()
}