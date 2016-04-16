package server

// Point = X Y Z coord
type Point struct {
	x int
	y int
	z int
}

// Direction = The way the snake is heading
type Direction string

// 6 Directions (6 god s/o Drake)
const (
	NORTH Direction = "North"
	EAST  Direction = "East"
	SOUTH Direction = "South"
	WEST  Direction = "West"
	UP    Direction = "Up"
	DOWN  Direction = "Down"
)

func move(p Point, d Direction) Point {
	switch d {
	case NORTH:
		return Point{p.x, p.y + 1, p.z}
	case EAST:
		return Point{p.x + 1, p.y, p.z}
	case SOUTH:
		return Point{p.x, p.y - 1, p.z}
	case WEST:
		return Point{p.x - 1, p.y, p.z}
	case UP:
		return Point{p.x, p.y, p.z + 1}
	case DOWN:
		return Point{p.x, p.y, p.z - 1}
	}
	panic("Invalid direction")
}

func (p Point) equals(other Point) bool {
	return (p.x == other.x) && (p.y == other.y) &&
		(p.z == other.z)
}
