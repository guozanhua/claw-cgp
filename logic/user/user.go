// Author: sheppard(ysf1026@gmail.com) 2013-12-25

package user

import (
	"fmt"
)

type User struct {
	Name string
	RecvMsg <-chan string
	SendMsg chan<- string
	Offline <-chan bool

	disconnected bool
}

func NewUser(name string, recv <-chan string, send chan<- string, offline <-chan bool) *User {
	return &User{name, recv, send, offline, false}
}

func (u *User) Tick() {
	for {
		select {
		case msg, ok := <-u.RecvMsg:
			if !ok {
				u.Logout()
				return
			}
			fmt.Println("UserTick", msg, Manager.users)
			Manager.broadcast <- msg
		case _, ok := <-u.Offline:
			if !ok {
				u.Logout()
			}
			u.Logout()
		}
	}
}

func (u *User) Send(msg string) {
	if !u.disconnected {
		u.SendMsg <- msg
	}
}

func (u *User) Logout() {
	if !u.disconnected {
		u.disconnected = true
		close(u.SendMsg)
		fmt.Println("User disconneted")
	}
}
