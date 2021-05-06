package server

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Eldius/poc-chat-websockets-go/chat"
	"golang.org/x/net/websocket"
)

func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func Start(port int) {
	server := chat.NewChatServer()
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.Handle("/echo", websocket.Handler(server.Accept))
	url := fmt.Sprintf(":%d", port)
	log.Printf("WS listening at: %s", url)
	err := http.ListenAndServe(url, mux)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
