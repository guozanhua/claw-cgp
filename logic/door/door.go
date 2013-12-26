// Author: sheppard(ysf1026@gmail.com) 2013-12-13

package door

import (
	"code.google.com/p/go.net/websocket"
	"github.com/yangsf5/claw-cgp/logic/user"
)

func init() {
}

func Login(conn *websocket.Conn, userName string) {
	offline := make(chan bool)

	recvMsg := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(conn, &msg)
			if err != nil {
				close(recvMsg)
				return
			}
			recvMsg <- msg
		}
	}()

	sendMsg := make(chan string)

	user.Manager.AddUser(user.NewUser(userName, recvMsg, sendMsg, offline))

	for {
		select {
		case msg, ok := <-sendMsg:
			// If the channel is closed, they disconnected.
			if !ok {
				return
			}

			if websocket.Message.Send(conn, &msg) != nil {
				// Disconneted.
				offline <- true
				return
			}
		}
	}
}

