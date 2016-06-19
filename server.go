package main

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
	log.Println("Listening localhost:3000..")
}
