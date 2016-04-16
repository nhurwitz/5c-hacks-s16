package server

type point struct {
	x int
	y int
	z int
}

// Direction = The way the snake is heading
type Direction string

// 6 Directions
const (
	NORTH Direction = "North"
	EAST  Direction = "East"
	SOUTH Direction = "South"
	WEST  Direction = "West"
	UP    Direction = "Up"
	DOWN  Direction = "Down"
)

func move(p point, d Direction) {
	switch d {
	case NORTH:
		p.y++
	case EAST:
		p.x++
	case SOUTH:
		p.y--
	case WEST:
		p.x--
	case UP:
		p.z++
	case DOWN:
		p.z--
	}
}
