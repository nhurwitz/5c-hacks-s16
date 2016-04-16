package server

// Snake = ~~~:)
type Snake struct {
	Head      Point     `json:"head"`
	Tail      []Point   `json:"tail"`
	Direction Direction `json:"direction"`
	ID        string    `json:"id"`
}

func NewSnake(sideLength int) Snake {
	return Snake{
		// ID:        uuid.NewV4().String(),
		Head:      randomPointIn(sideLength),
		Tail:      make([]Point, 0),
		Direction: randomDirection(),
	}
}

func (snake Snake) tickedHead() Point {
	return snake.move(false).Head
}

// 0-indexed grid. (i.e., if sidelength = 4, [3] = edge)
func (snake Snake) collidedWithEdge(sideLength int) bool {
	return (snake.Head.X >= sideLength) || (snake.Head.X < 0) ||
		(snake.Head.Y >= sideLength) || (snake.Head.Y < 0) ||
		(snake.Head.Z >= sideLength) || (snake.Head.Z < 0)
}

func (snake Snake) collidedWithSelf() bool {
	for _, el := range snake.Tail {
		return snake.Head.equals(el)
	}
	return false
}

func (snake Snake) collidedWithOther(other Snake) bool {
	if snake.Head.equals(other.Head) {
		return true
	}
	for _, otherPoint := range other.Tail {
		if snake.Head.equals(otherPoint) {
			return true
		}
	}
	return false
}

func (snake Snake) containsPoint(p Point) bool {
	for _, bodyPoint := range append(snake.Tail, snake.Head) {
		if bodyPoint == p {
			return true
		}
	}

	return false
}

// Advance snake by adding head to the tail
// Remove last element if not capturing
func (snake Snake) move(capturing bool) Snake {
	var newHead = move(snake.Head, snake.Direction)
	snake.Tail = append([]Point{snake.Head}, snake.Tail...)
	snake.Head = newHead
	if !capturing {
		snake.Tail = snake.Tail[:len(snake.Tail)-1]
	}
	return snake
}
