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

func (w World) requeuePoints() World {
	for len(w.pendingPoints) < len(w.snakes)*POINT_RATIO {
		newPoint := w.randomPoint()
		// TODO XXX don't generate points currently in snakes?
		for w.isPending(newPoint) {
			newPoint = w.randomPoint()
		}
		w.pendingPoints = append(w.pendingPoints, newPoint)
	}

	return w
}

func (w World) isPending(p Point) bool {
	for _, point := range w.pendingPoints {
		if p.equals(point) {
			return true
		}
	}
	return false
}

func Tick(w World) (World, []Event) {

	// Tick heads. Map of snake IDs to new head
	tickedHeads := make(map[string]Point)
	for _, snake := range w.snakes {
		tickedHeads[snake.id] = snake.tickedHead()
	}

	// Which snakes are capturing? If a snake ID is present, the snake is
	// capturing.
	snakesWhichAreCapturing := make(map[string]bool)
	for snakeID, newHead := range tickedHeads {
		if w.isPending(newHead) {
			snakesWhichAreCapturing[snakeID] = true
		}
	}

	// Move each snake, note whether or not capturing. We'll filter these by
	// whether or not they're still alive next.
	livingMovedSnakes := make(map[string]Snake)
	for _, snake := range w.snakes {
		livingMovedSnakes[snake.id] = snake.move(snakesWhichAreCapturing[snake.id])
	}

	// Collision detection. DO NOT REMOVE FROM THE MAP YET; JUST ASSEMBLE IDs.
	deadSnakeIDs := make(map[string]bool)
	for snakeID, snake := range livingMovedSnakes {

		// colliding with others.
		for collidedSnakeID, collidingSnake := range livingMovedSnakes {
			areSameSnake := snakeID == collidedSnakeID
			if snake.collidedWithOther(collidingSnake) && !areSameSnake {
				deadSnakeIDs[snakeID] = true
			}
		}

		// colliding with self.
		if snake.collidedWithSelf() {
			deadSnakeIDs[snakeID] = true
		}

		// colliding with edge.
		if snake.collidedWithEdge(w.gridLength) {
			deadSnakeIDs[snakeID] = true
		}
	}

	// Remove dead snakes
	for snakeID := range deadSnakeIDs {
		delete(livingMovedSnakes, snakeID)
	}

	// Track events - who died?
	var events []Event
	for snakeID := range deadSnakeIDs {
		events = append(events, Event{
			eventType: EVENT_DIE,
			snakeId:   snakeID,
		})
	}

	// Regenerate pending points
	w = w.requeuePoints()

	// Update world. Return the world + events
	var newSnakes []Snake
	for _, snake := range livingMovedSnakes {
		newSnakes = append(newSnakes, snake)
	}
	w.snakes = newSnakes

	return w, events
}
