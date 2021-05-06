package chat

import (
	"log"

	"golang.org/x/net/websocket"
)

type Message struct {
	Msg  string
	Name string
	Room string
}

type ChatServer struct {
	clients []*websocket.Conn
}

func NewChatServer() *ChatServer {
	return &ChatServer{}
}

func (c *ChatServer) Accept(ws *websocket.Conn) {
	c.clients = append(c.clients, ws)
	websocket.JSON.Send(ws, Message{
		Msg: "Test!",
	})
	c.listen(ws)
}

func (c *ChatServer) RemoveClient(ws *websocket.Conn) {
	aux := make([]*websocket.Conn, 0)
	for i, _ws := range c.clients {
		if ws != _ws {
			aux = append(aux, _ws)
		} else {
			log.Printf("Removing socket '%d'\n", i)
		}
	}
	c.clients = aux
}

func (c *ChatServer) broadcast(msg Message, ws *websocket.Conn) {
	for _, _ws := range c.clients {
		if err := websocket.JSON.Send(_ws, msg); err != nil {
			log.Printf("can't marshal message: %s\n", err)
		} else {
			log.Println("Message sent")
		}
	}
}

func (c *ChatServer) listen(ws *websocket.Conn) {
	for {
		var msg Message
		if err := websocket.JSON.Receive(ws, &msg); err != nil {
			log.Printf("Failed to read message: '%s'", err.Error())
			if err.Error() == "EOF" {
				c.RemoveClient(ws)
				return
			}
		} else {
			log.Printf("received message from '%s'/'%s': '%v'", ws.RemoteAddr().String(), ws.LocalAddr().String(), msg)
			c.broadcast(msg, ws)
		}
		//var msgb []byte
		//if err := websocket.Message.Receive(ws, &msgb); err != nil {
		//	log.Panicf("Failed to read message: '%s'", err.Error())
		//}
		//log.Printf("received message from '%s'/'%s': '%s'", ws.RemoteAddr().String(), ws.LocalAddr().String(), string(msgb))
		//if err := json.Unmarshal(msgb, &msg); err != nil {
		//	log.Printf("Failed to unmarshal message: '%s'\n", err.Error())
		//} else {
		//	log.Printf("received message from '%s'/'%s': '%v'", ws.RemoteAddr().String(), ws.LocalAddr().String(), msg)
		//	c.broadcast(msg, ws)
		//}
	}
}
