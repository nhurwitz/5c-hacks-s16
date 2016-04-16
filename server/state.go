package server

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	world      = newWorld(20)
	eventChan  = make(chan Event, 10000)
	actionChan = make(chan Action, 10000)

	connections = make(map[*websocket.Conn]bool)
	mu          = new(sync.Mutex)

	rate = flag.Int("rate", 250, "num ms between world ticks")
)

func init() {

	rand.Seed(time.Now().UnixNano())

	flag.Parse()

	//actionChan <- Action{
	//ActionType: ActionSpawn,
	//SnakeID:    uuid.NewV4().String(),
	//}

	ticker := time.NewTicker(time.Duration((*rate)) * time.Millisecond).C
	go func() {
		for {
			var evts []Event
			select {
			case <-ticker: // world changes
				world, evts = Tick(world)
				eventChan <- Event{
					EventType: EventWorld,
					World:     world,
				}
				fmt.Println(world)
			case a := <-actionChan: // update the world in response to an action
				world, evts = Act(world, a)
			case e := <-eventChan:
				mu.Lock()
				for conn := range connections {
					// ignore the error because it's 1:17 and we on this thang
					if err := conn.WriteJSON(e); err != nil {
						fmt.Println("error writing to global connection:", err)
					}
				}
				mu.Unlock()
			}

			for _, evt := range evts {
				eventChan <- evt
			}
		}
	}()

}
