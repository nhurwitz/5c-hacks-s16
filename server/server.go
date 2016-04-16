package server

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", GameHandler)
	http.ListenAndServe(":8000", r)
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
}

func startServer() {

}
