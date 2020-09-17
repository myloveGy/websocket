package main

import (
	"fmt"
	"net/http"
	"websocket/api"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// socket 处理
	http.HandleFunc("/ws", api.WebSocket)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("listen: 3000")
	_ = http.ListenAndServe(":3000", nil)
}
