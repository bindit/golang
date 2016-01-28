package object

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