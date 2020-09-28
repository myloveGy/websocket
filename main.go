package main

import (
	"fmt"
	"net/http"
	"websocket/api"
)

func main() {

	// socket 处理
	http.HandleFunc("/ws", api.WebSocket)

	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	fmt.Println("listen: 3000")
	_ = http.ListenAndServe(":3000", nil)
}
