package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"websocket/api"
	"websocket/api/ws"
	"websocket/utils"
)

func main() {

	hub := ws.NewHub()

	go hub.Run()

	http.HandleFunc("/info", func(w http.ResponseWriter, request *http.Request) {
		info := map[string]interface{}{
			"clients": len(hub.Clients),
			"time":    utils.DateTime(),
		}

		str, _ := json.Marshal(info)
		w.Write(str)
	})

	// socket 处理
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		api.WebSocket(hub, w, r)
	})

	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	fmt.Println("listen: 3000")
	_ = http.ListenAndServe(":3000", nil)
}
