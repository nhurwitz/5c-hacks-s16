package server

import "time"

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
			case <-timer: // world changes
				w, evts = Tick(w)
				eventChan <- Event{
					EventType: EventWorld,
					World:     w,
				}
			case a := <-actionChan: // update the world in response to an action
				w, evts = Act(w, a)
			}

			for _, evt := range evts {
				eventChan <- evt
			}
		}
	}()

}
