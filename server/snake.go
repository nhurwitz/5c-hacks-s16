package server

// Snake = ~~~:)
type Snake struct {
	head      Point
	tail      []Point
	direction Direction
	id        string
}

func (snake Snake) collidedWithEdge(sideLength int) bool {
	return false
}

func (snake Snake) collidedWithSelf() bool {
	return false
}

func (snake Snake) collidedWithOther(other Snake) bool {
	return false
}

func (snake Snake) move(capturing bool) Snake {
	return snake
}
