//package in go.mod is called chat.com
// should I add package server or chat.com/server?

package main

import (
	"github.com/gorilla/websocket"
	// "log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	connection.Close()
}

func main() {

}
