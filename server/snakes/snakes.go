package main

import "github.com/nhurwitz/5c-hacks-s16/server"

// For performance monitoring.
import _ "net/http/pprof"

func main() {
	server.StartServer()
}
