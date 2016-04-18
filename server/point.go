package server

import "math/rand"

// Point = X Y Z coord
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

// Direction = The way the snake is heading
type Direction string

// 6 Directions (6 god s/o Drake)
const (
	North Direction = "North" // 1
	East  Direction = "East"  // 2
	South Direction = "South" // 3
	West  Direction = "West"  // 4
	Up    Direction = "Up"    // 5
	Down  Direction = "Down"  // 6
)

func move(p Point, d Direction, sideLength int) Point {
	switch d {
	case North:
		return Point{p.X, (p.Y + 1) % sideLength, p.Z}
	case East:
		return Point{(p.X + 1) % sideLength, p.Y, p.Z}
	case Up:
		return Point{p.X, p.Y, (p.Z + 1) % sideLength}
	case West:
		if p.X == 0 {
			return Point{sideLength - 1, p.Y, p.Z}
		}
		return Point{p.X - 1, p.Y, p.Z}
	case South:
		if p.Y == 0 {
			return Point{p.X, sideLength - 1, p.Z}
		}
		return Point{p.X, p.Y - 1, p.Z}
	case Down:
		if p.Z == 0 {
			return Point{p.X, p.Y, p.Z - 1}
		}
		return Point{p.X, p.Y, sideLength - 1}
	}
	panic("Invalid direction")
}

func randomPointIn(sideLength int) Point {
	return Point{
		X: rand.Intn(sideLength),
		Y: rand.Intn(sideLength),
		Z: rand.Intn(sideLength),
	}
}

func randomDirection() Direction {
	directions := []Direction{North, East, South, West, Down, Up}
	return directions[rand.Intn(len(directions))]
}
