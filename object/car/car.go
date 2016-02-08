package car

type Rider interface {
	Speed() int
}

type Speeder interface {
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