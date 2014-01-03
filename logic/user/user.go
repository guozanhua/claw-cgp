// Author: sheppard(ysf1026@gmail.com) 2013-12-25

package user

import (
	"fmt"
)

type User struct {
	Name string
	RecvMsg <-chan string
	SendMsg chan<- string
	Offline <-chan error

	disconnected bool
}

func NewUser(name string, recv <-chan string, send chan<- string, offline <-chan error) *User {
	return &User{name, recv, send, offline, false}
}

func (u *User) Tick() {
	for {
		select {
		case msg, ok := <-u.RecvMsg:
			if !ok {
				u.Logout("Recv channel closed")
				return
			}
			fmt.Println("UserTick", msg, Manager.users)
			Manager.broadcast <- msg
		case err, ok := <-u.Offline:
			if !ok {
				u.Logout("Offline channel closed")
			} else {
				u.Logout(err.Error())
			}
		}
	}
}

func (u *User) Send(msg string) {
	if !u.disconnected {
		u.SendMsg <- msg
	}
}

func (u *User) Logout(reason string) {
	if !u.disconnected {
		u.disconnected = true
		close(u.SendMsg)
		fmt.Println("User disconneted, err:", reason)
	}
}
