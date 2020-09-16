package main

import (
	"fmt"
	"log"
	"net/http"
	"websocket/utils"

	"github.com/gorilla/websocket"
)

// Message 发放的消息信息
type Message struct {
	Type   string `json:"type"`
	Data   string `json:"data"`
	Source string `json:"source"`
}

var upgrade = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, err = upgrade.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}

		defer conn.Close()

		for {
			msg := &Message{}
			err := conn.ReadJSON(msg)
			if err != nil {
				log.Println("read: ", err)
				break
			}

			log.Printf("recv: %v", msg)
			result, err := utils.GetHTTP(msg.Data)
			if err != nil || result.Code != 0 {
				log.Printf("机器人回复失败：%v\n", err)
				break
			}

			log.Println(result)
			err = conn.WriteJSON(&Message{
				Type:   "system",
				Data:   result.Content,
				Source: "system",
			})
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("listen: 3000")
	_ = http.ListenAndServe(":3000", nil)
}
