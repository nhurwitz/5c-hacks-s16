package server

type ActionEnum string

// Possible Actions
const (
	ActionSpawn           ActionEnum = "Spawn"     // snake spawns
	ActionChangeDirection ActionEnum = "Direction" // snake changes direction
	ActionQuit            ActionEnum = "Quit"
)

type Action struct {
	ActionType ActionEnum `json:"actionType"`
	SnakeID    string     `json:"snakeID"`
	Direction  *Direction `json:"direction,omitempty"`
}
