package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Client struct {
	ws    *websocket.Conn
	send  chan []byte
	world chan string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ayy")
}

func adrilonHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	client := Client{
		ws:    conn,
		send:  make(chan []byte),
		world: make(chan string),
	}

	clients <- client
}

var upgrader = websocket.Upgrader{}
var clients = make(chan Client)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/Adrilon", adrilonHandler)
	http.ListenAndServe(":8080", nil)
}
