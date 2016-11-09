package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
)

type Client struct {
	ws    *websocket.Conn
	send  chan []byte
	world chan string
}

// Index Handler
func indexHandler(w http.ResponseWriter, r *http.Request) {

	type Page struct {
	}

	p := &Page{}
	t, err := template.ParseFiles("pages/index.html")

	if err == nil {
		log.Print("Writing index.html out")
		t.Execute(w, p)
	} else {
		log.Fatal("Cannot find pages/index.html")
		fmt.Fprintf(w, "Failure")
	}
}

func adrilonHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	log.Print("Creating new client")
	client := Client{
		ws:    conn,
		send:  make(chan []byte),
		world: make(chan string),
	}

	log.Print("Adding client to world")
	world.clients <- client
}

// Globals... can we avoid this?

var world = Zone{
	clients: make(chan Client),
}

var upgrader = websocket.Upgrader{}

func main() {

	fmt.Println("whatever")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/Adrilon", adrilonHandler)

	fmt.Println("Start server")
	http.ListenAndServe(":8080", nil)
}
