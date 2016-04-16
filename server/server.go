package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", SocketHandler)
	r.HandleFunc("/test", EchoHandler)
	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", r)
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world!"))
	w.WriteHeader(200)
}

// start a goroutine which, for every event, pulls that event off and writes it
// to each connections

func SocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// Add this connection to the
	mu.Lock()
	connections[conn] = true
	mu.Unlock()

	snakeID := uuid.NewV4().String()

	// upon connecting, add a snake, broadcast an event and shit and send a
	// message to the client saying which snake it is
	actionChan <- Action{
		ActionType: ActionSpawn,
		SnakeID:    snakeID,
	}

	// message the client, and this client only, with its ID and the world
	err = conn.WriteJSON(Event{
		EventType: EventWelcome,
		SnakeID:   &snakeID,
		World:     world,
	})
	if err != nil {
		fmt.Println("ERROR SENDING WELCOME EVENT:", err)
	}

	// spin up a goroutine which, for as long as the client is connected, pulls
	// actions off of the websocket thing and pushes them onto the actionsChan.
	// when this closes, put action for snake leaving on action chan.
	go func(conn *websocket.Conn) {
		for {

			var a Action
			if err := conn.ReadJSON(&a); err != nil {
				actionChan <- Action{
					ActionType: ActionQuit,
					SnakeID:    snakeID,
				}

				// remove the websocket connection from the pool of connections
				mu.Lock()
				delete(connections, conn)
				mu.Unlock()

				return // this will terminate the goroutine
			}

			a.SnakeID = snakeID
			actionChan <- a

		}
	}(conn)

}
