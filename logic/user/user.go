// Author: sheppard(ysf1026@gmail.com) 2013-12-25

package user

import (
	"fmt"
	"github.com/yangsf5/claw-cgp/logic/proto"
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
			fmt.Println("User recv:", msg)
			chatMsg := &proto.HCChat{u.Name, msg}
			Manager.broadcast <- proto.Encode(chatMsg)
		case err, ok := <-u.Offline:
			if !ok {
				u.Logout("Offline channel closed")
			} else {
				u.Logout(err.Error())
			}
			return
		}
	}
}

func (u *User) Send(msg string) {
	if !u.disconnected {
		u.SendMsg <- msg
		fmt.Println("User send:", msg)
	}
}

func (u *User) Login() {
	roomMsg := &proto.HCRoomList{[]proto.Room{{"chess", 12}, {"poker", 100}}}
	fmt.Println(roomMsg, u.disconnected)
	u.Send(proto.Encode(roomMsg))
	fmt.Println("User login, name:", u.Name)
}

func (u *User) Logout(reason string) {
	if !u.disconnected {
		u.disconnected = true
		close(u.SendMsg)
		Manager.DelUser(u)
		fmt.Println("User disconneted, err:", reason)
	}
}
