package server

type EventEnum string

const (
	EVENT_SPAWN EventEnum = "Spawn"
	EVENT_DIE   EventEnum = "Die"
	EVENT_WORLD EventEnum = "World"
	EVENT_JOIN  EventEnum = "Join"
	EVENT_LEAVE EventEnum = "Leave"
)

type Event struct {
	eventType EventEnum
	snakeId   string
	world     World
}
