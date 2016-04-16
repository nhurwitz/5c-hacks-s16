package server

import "time"

// This file contains global state - and World, protected by a mutex.

var (
	w          = newWorld(20)
	eventChan  = make(chan Event)
	actionChan = make(chan Action)
)

func init() {

	timer := time.NewTimer(250 * time.Millisecond).C
	go func() {
		for {
			var evts []Event
			select {
			case <-timer:
				w, evts = Tick(w)
			case a := <-actionChan:
				w, evts = Act(w, a)
			}

			for _, evt := range evts {
				eventChan <- evt
			}
		}
	}()

}
