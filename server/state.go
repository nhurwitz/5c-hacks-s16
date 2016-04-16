package server

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	world      = newWorld(20)
	eventChan  = make(chan Event)
	actionChan = make(chan Action)

	connections map[*websocket.Conn]bool
	mu          = new(sync.Mutex)
)

func init() {

	timer := time.NewTimer(250 * time.Millisecond).C
	go func() {
		for {
			var evts []Event
			select {
			case <-timer: // world changes
				world, evts = Tick(world)
				eventChan <- Event{
					EventType: EventWorld,
					World:     world,
				}
			case a := <-actionChan: // update the world in response to an action
				world, evts = Act(world, a)
			case e := <-eventChan:
				mu.Lock()
				for conn := range connections {
					// ignore the error because it's 1:17 and we on this thang
					conn.WriteJSON(e)
				}
				mu.Unlock()
			}

			for _, evt := range evts {
				eventChan <- evt
			}
		}
	}()

}
