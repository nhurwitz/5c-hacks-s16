package server

// Point = X Y Z coord
type Point struct {
	X int
	Y int
	Z int
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
		return Point{p.X, p.Y + 1, p.Z}
	case EAST:
		return Point{p.X + 1, p.Y, p.Z}
	case SOUTH:
		return Point{p.X, p.Y - 1, p.Z}
	case WEST:
		return Point{p.X - 1, p.Y, p.Z}
	case UP:
		return Point{p.X, p.Y, p.Z + 1}
	case DOWN:
		return Point{p.X, p.Y, p.Z - 1}
	}
	panic("Invalid direction")
}

func (p Point) equals(other Point) bool {
	return (p.X == other.X) && (p.Y == other.Y) &&
		(p.Z == other.Z)
}
