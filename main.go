package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有连接
		return true
	},
}

func main() {
	http.HandleFunc("/", handleWebSocket)

	fmt.Println("WebSocket server started at ws://localhost:5000/")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	//
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("recv: %s", message)
		if err = conn.WriteMessage(mt, message); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
