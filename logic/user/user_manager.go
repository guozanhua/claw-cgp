// Author: sheppard(ysf1026@gmail.com) 2013-12-25

package user

import (
	"fmt"
)

type UserManager struct {
	users map[string] *User
	broadcast chan string
}

func NewUserManager() *UserManager {
	m := new(UserManager)
	m.users = make(map[string] *User)
	m.broadcast = make(chan string)
	return m
}

func (m *UserManager) AddUser(user *User) bool {
	if _, ok := m.users[user.Name]; ok {
		return false;
	}
	m.users[user.Name] = user;
	go user.Tick()
	return true;
}

func (m *UserManager) DelUser(user *User) {
	delete(m.users, user.Name)
}

func (m *UserManager) Tick() {
	for {
		msg := <-m.broadcast
		fmt.Println("M Tick", msg)
		for _, user := range m.users {
			user.Send(msg)
		}
	}
}

var Manager *UserManager

func init() {
	Manager = NewUserManager()

	go Manager.Tick()
}

