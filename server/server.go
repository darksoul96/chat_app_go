//package in go.mod is called chat.com
// should I add package server or chat.com/server?

package main

import (
	// "github.com/gorilla/websocket"
	// "log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// conn, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	log.Printf("Upgrade error: %v", err)
	// 	return
	// }

	// defer conn.Close()
	// for {
	// 	mt, message, err := conn.ReadMessage()
	// 	if err != nil {
	// 		log.Printf("Read error: %v", err)
	// 		return
	// 	}
	// 	log.Printf("Received: %s", message)
	// 	err = conn.WriteMessage(mt, message)
	// 	if err != nil {
	// 		log.Printf("Write error: %v", err)
	// 		return
	// 	}
	// }
}

func main() {
	// var upgrader = websocket.Upgrader{
	// 	ReadBufferSize:  1024,
	// 	WriteBufferSize: 1024,
	// }

}
