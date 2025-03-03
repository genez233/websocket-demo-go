package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

// 定义一个 upgrader，用于将 HTTP 连接升级为 WebSocket 连接
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有连接
		return true
	},
}

func main() {
	// 设置 WebSocket 路由
	http.HandleFunc("/", handleWebSocket)

	// 启动 HTTP 服务器
	fmt.Println("WebSocket server started at ws://localhost:5000/")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

// 处理 WebSocket 连接
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// 循环读取客户端发送的消息
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// 打印接收到的消息
		log.Printf("recv: %s", message)

		// 将消息回显给客户端
		if err = conn.WriteMessage(mt, message); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
