// Author: sheppard(ysf1026@gmail.com) 2014-01-12

package hall

import (
	"fmt"
)

type User interface {
	Send(string)
}

var (
	users map[string] User

	broadcast chan string
)

func init() {
	users = make(map[string] User)
	broadcast = make(chan string)

	go Tick()
}

func AddUser(uid string, u User) bool {
	if _, ok := users[uid]; ok {
		return false
	}
	users[uid] = u
	return true
}

func DelUser(uid string) {
	delete(users, uid)
}

func Broadcast(msg string) {
	broadcast <- msg
}

func Tick() {
	for {
		msg := <-broadcast
		fmt.Println("Hall broadcast msg:", msg)
		for _, u := range users {
			u.Send(msg)
		}
	}
}

