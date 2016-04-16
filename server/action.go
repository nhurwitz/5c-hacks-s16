package server

type ActionEnum string

// Possible Actions
const (
	ACTION_SPAWN     ActionEnum = "Spawn"     // snake spawns
	ACTION_DIRECTION ActionEnum = "Direction" // snake changes direction
	ACTION_QUIT      ActionEnum = "Quit"      // snake quits
)

type Action struct {
	ActionType ActionEnum `json:"actionType"`
	SnakeID    string     `json:"snakeID"`
	Direction  *Direction `json:"direction,omitempty"`
}
