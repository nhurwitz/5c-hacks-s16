package server

// "github.com/satori/go.uuid"
// "math/rand"

const PointRatio = 1

type World struct {
	SideLength    int              `json:"sideLength"`
	PendingPoints []Point          `json:"pendingPoints"`
	Snakes        map[string]Snake `json:"snakes"`
}

func newWorld(gridLength int) World {
	return World{
		SideLength:    gridLength,
		PendingPoints: make([]Point, 0),
		Snakes:        make(map[string]Snake),
	}
}

func (w World) randomPoint() Point {
	return randomPointIn(w.SideLength)
}

func (w World) anySnakesContain(p Point) bool {
	for _, snake := range w.Snakes {
		if snake.containsPoint(p) {
			return true
		}
	}

	return false
}

func (w World) pointInUse(p Point) bool {
	if w.anySnakesContain(p) {
		return true
	}

	for _, pending := range w.PendingPoints {
		if pending == p {
			return true
		}
	}

	return false
}

func (w World) requeuePoints() World {

	for len(w.PendingPoints) < len(w.Snakes)*PointRatio {
		newPoint := w.randomPoint()
		pending, _ := w.isPending(newPoint)
		for pending || w.anySnakesContain(newPoint) {
			newPoint = w.randomPoint()
			pending, _ = w.isPending(newPoint)
		}
		w.PendingPoints = append(w.PendingPoints, newPoint)
	}

	return w
}

func (w World) isPending(p Point) (bool, int) {
	for i, point := range w.PendingPoints {
		if p.equals(point) {
			return true, i
		}
	}
	return false, -1
}

func Tick(w World) (World, []Event) {

	// Tick heads. Map of snake IDs to new head
	tickedHeads := make(map[string]Point)
	for _, snake := range w.Snakes {
		tickedHeads[snake.ID] = snake.tickedHead()
	}

	// Which snakes are capturing? If a snake ID is present, the snake is
	// capturing.
	// XXX TODO remove pending points
	snakesWhichAreCapturing := make(map[string]bool)
	for snakeID, newHead := range tickedHeads {
		if pending, i := w.isPending(newHead); pending {
			snakesWhichAreCapturing[snakeID] = true
			w.PendingPoints = append(w.PendingPoints[:i], w.PendingPoints[i+1:]...)
		}
	}

	// Move each snake, note whether or not capturing. We'll filter these by
	// whether or not they're still alive next.
	livingMovedSnakes := make(map[string]Snake)
	for _, snake := range w.Snakes {
		livingMovedSnakes[snake.ID] = snake.move(snakesWhichAreCapturing[snake.ID])
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
		if snake.collidedWithEdge(w.SideLength) {
			deadSnakeIDs[snakeID] = true
		}
	}

	// Remove dead snakes
	for snakeID := range deadSnakeIDs {
		delete(livingMovedSnakes, snakeID)

		// XXX remove this
		actionChan <- Action{
			ActionType: ActionSpawn,
			SnakeID:    snakeID,
		}
	}

	// Track events - who died?
	var events []Event
	for snakeID := range deadSnakeIDs {
		events = append(events, Event{
			EventType: EventDie,
			SnakeID:   &snakeID,
		})
	}

	// Regenerate pending points
	w = w.requeuePoints()

	// Update world. Return the world + events
	w.Snakes = livingMovedSnakes

	return w, events
}

// Act on user action
func Act(w World, a Action) (World, []Event) {

	switch a.ActionType {

	case ActionChangeDirection:
		var temp Snake
		temp = w.Snakes[a.SnakeID]
		temp.Direction = *a.Direction
		w.Snakes[a.SnakeID] = temp
		return w, nil

	case ActionSpawn:
		newSnake := NewSnake(w.SideLength)
		newSnake.ID = a.SnakeID // we need it to be the same player

		// Makes sure new head contained within another snake / a pending point
	validationLoop:
		for {
			if w.pointInUse(newSnake.Head) {
				newSnake.Head = randomPointIn(w.SideLength)
				continue validationLoop
			}

			break
		}

		w.Snakes[newSnake.ID] = newSnake
		eventArr := []Event{Event{
			EventType: EventSpawn,
			SnakeID:   &newSnake.ID,
		}}

		return w, eventArr

	case ActionQuit:
		delete(w.Snakes, a.SnakeID)
		w.PendingPoints = w.PendingPoints[1:]
		return w, []Event{{
			EventType: EventLeave,
			SnakeID:   &a.SnakeID,
		}}

	default:
		panic("invalid action enum")
	}

	return w, nil
}
