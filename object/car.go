package object

import "math/rand"

type Rider interface {
	Speed() int
}

type Car struct {
	Color string
	Engine Engine
	Door Door
}

type Door struct {

}

type Engine struct {
	Power int
	Capacity float32
}

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

type Coupe struct {
	Car
}

func (c Coupe) Speed() int {
	return 100
}