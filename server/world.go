package server

import "math/rand"

const POINT_RATIO = 1

type World struct {
	gridLength    int
	pendingPoints []Point
	snakes        []Snake
}

func newWorld(gridLength int) World {
	return World{
		gridLength:    gridLength,
		pendingPoints: make([]Point, 0),
		snakes:        make([]Snake, 0)}
}

func (w World) randomPoint() Point {
	return Point{
		x: rand.Intn(w.gridLength),
		y: rand.Intn(w.gridLength),
		z: rand.Intn(w.gridLength)}
}

func (w World) requeuePoints() {
	for len(w.pendingPoints) < len(w.snakes)*POINT_RATIO {
		newPoint := w.randomPoint()
		// TODO XXX don't generate points currently in snakes?
		for w.pointContains(newPoint) {
			newPoint = w.randomPoint()
		}
		w.pendingPoints = append(w.pendingPoints, newPoint)
	}
}

func (w World) pointContains(p Point) bool {
	for _, point := range w.pendingPoints {
		if p.equals(point) {
			return true
		}
	}
	return false
}
