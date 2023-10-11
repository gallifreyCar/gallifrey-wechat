package server

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn      *websocket.Conn
	Name      string
	DataQueue chan []byte
}

func (c *Client) Read() {
	defer func() {
		MyServer.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// PongHandler() 是一个处理器，用于处理从对等方接收到的Pong消息。
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			MyServer.Unregister <- c
			c.Conn.Close()
			break
		}
		if string(message) == "ping" {
			MyServer.Broadcast <- []byte("pong")
			continue
		}
		if string(message) == "test" {
			targetClient := MyServer.Clients["test01"]
			targetClient.DataQueue <- []byte("test single send")
			continue
		}

		// 将消息发送到广播通道
		MyServer.Broadcast <- message

	}
}

func (c *Client) Write() {
	defer func() {
		MyServer.Unregister <- c
		c.Conn.Close()
	}()
	for {

		select {
		case message := <-c.DataQueue:
			err := c.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				MyServer.Unregister <- c
				c.Conn.Close()
				break
			}
		case message := <-MyServer.Broadcast:
			err := c.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				MyServer.Unregister <- c
				c.Conn.Close()
				break
			}
		}
	}
}
