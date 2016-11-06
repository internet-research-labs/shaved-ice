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

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/Adrilon", indexHandler)
	http.ListenAndServe(":8080", nil)
}
