package server

type EventEnum string

const (
	EventSpawn   EventEnum = "Spawn"
	EventDie     EventEnum = "Die"
	EventWorld   EventEnum = "World"
	EventJoin    EventEnum = "Join"
	EventLeave   EventEnum = "Leave"
	EventWelcome EventEnum = "Welcome"
)

type Event struct {
	EventType EventEnum `json:"eventType"`
	SnakeID   *string   `json:"snakeID,omitempty"`
	World     World     `json:"world"`
}
