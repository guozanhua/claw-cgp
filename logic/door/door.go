// Author: sheppard(ysf1026@gmail.com) 2013-12-13

package door

import (
	"fmt"
	"code.google.com/p/go.net/websocket"
)

func init() {
}

func Login(conn *websocket.Conn) {
	newMessage := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(conn, &msg)
			if err != nil {
				close(newMessage)
				return
			}
			newMessage <- msg
		}
	}()

	for {
		select {
		case msg, ok := <-newMessage:
			// If the channel is closed, they disconnected.
			if !ok {
				return
			}
			fmt.Println("New message", msg)
			//TODO
		}
	}
}

