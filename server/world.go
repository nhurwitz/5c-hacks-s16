package server

import "math/rand"

const POINT_RATIO = 1

type World struct {
	gridLength      int
	availablePoints []Point
	snakes          []Snake
}

func newWorld(gridLength int) World {
	return World{
		gridLength:      gridLength,
		availablePoints: make([]Point, 0),
		snakes:          make([]Snake, 0)}
}

func (w World) randomPoint() Point {
	return Point{
		x: rand.Intn(w.gridLength),
		y: rand.Intn(w.gridLength),
		z: rand.Intn(w.gridLength)}
}

func (w World) requeuePoints() {
	for len(w.availablePoints) < len(w.snakes)*POINT_RATIO {
		newPoint := w.randomPoint()
		for w.pointContains(newPoint) {
			newPoint = w.randomPoint()
		}
		w.availablePoints = append(w.availablePoints, newPoint)
	}
}

func (w World) pointContains(p Point) bool {
	for _, point := range w.availablePoints {
		if p.equals(point) {
			return false
		}
	}
	return true
}
