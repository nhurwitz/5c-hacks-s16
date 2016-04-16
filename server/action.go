package server

type ActionEnum string

// Possible Actions
const (
	ActionSpawn     ActionEnum = "Spawn"     // snake spawns
	ActionDirection ActionEnum = "Direction" // snake changes direction
	ActionQuit      ActionEnum = "Quit"      // snake quits
)

type Action struct {
	ActionType ActionEnum `json:"actionType"`
	SnakeID    string     `json:"snakeID"`
	Direction  *Direction `json:"direction,omitempty"`
}
