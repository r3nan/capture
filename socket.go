package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

var socket socketio.Socket

func getSocketHandler() http.Handler {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		socket = so
		emit(captures)
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("socket error:", err)
	})
	return server
}

func emit(captures []Capture) {
	if socket == nil {
		return
	}
	socket.Emit("captures", captures)
}
