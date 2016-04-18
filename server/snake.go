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

func (snake Snake) tickedHead(sideLength int) Point {
	return snake.move(false, sideLength).Head
}

func (snake Snake) collidedWithSelf() bool {
	for _, pt := range snake.Tail {
		if pt == snake.Head {
			return true
		}
	}

	return false
}

func (snake Snake) collidedWithOther(other Snake) bool {
	if snake.Head == other.Head {
		return true
	}

	for _, otherPoint := range other.Tail {
		if snake.Head == otherPoint {
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
func (snake Snake) move(capturing bool, sideLength int) Snake {
	var newHead = move(snake.Head, snake.Direction, sideLength)
	snake.Tail = append([]Point{snake.Head}, snake.Tail...)
	snake.Head = newHead
	if !capturing {
		snake.Tail = snake.Tail[:len(snake.Tail)-1]
	}
	return snake
}
