package server

type ActionEnum string

// Possible Actions
const (
	SPAWN     ActionEnum = "Spawn"     // snake spawns
	DIRECTION ActionEnum = "Direction" // snake changes direction
	QUIT      ActionEnum = "Quit"      // snake quits
)

type Action struct {
	actionType ActionEnum
	snakeId    string
}
