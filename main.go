package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrade.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			ch := time.Tick(60 * time.Second)
			for range ch {
				mType, msg, _ := conn.ReadMessage()
				fmt.Printf("type: %v, message: %v \n", mType, string(msg))
				_ = conn.WriteMessage(mType, msg)
			}
		}(conn)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	_ = http.ListenAndServe(":3000", nil)
}
