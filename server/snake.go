package server

// Snake = ~~~:)
type Snake struct {
	head      Point
	tail      []Point
	direction Direction
	id        string
}

// 0-indexed grid. (i.e.) if sidelength = 4, 3 edge
func (snake Snake) collidedWithEdge(sideLength int) bool {
	return (snake.head.x >= sideLength) && (snake.head.x < 0) &&
		(snake.head.y >= sideLength) && (snake.head.y < 0) &&
		(snake.head.z >= sideLength) && (snake.head.z < 0)
}

func (snake Snake) collidedWithSelf() bool {
	for _, el := range snake.tail {
		return snake.head == el
	}
	return false
}

func (snake Snake) collidedWithOther(other Snake) bool {

}

func (snake Snake) move(capturing bool) Snake {
	return snake
}
