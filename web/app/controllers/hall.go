// Author: sheppard(ysf1026@gmail.com) 2013-12-04

package controllers

import (
	"code.google.com/p/go.net/websocket"
	"github.com/robfig/revel"
	"github.com/yangsf5/claw-cgp/logic/door"
)

type Hall struct {
	*revel.Controller
}

func (c Hall) Hall(user string) revel.Result {
	return c.Render(user)
}

func (c Hall) Socket(user string, ws *websocket.Conn) revel.Result {
	revel.INFO.Printf("%s", user)
	door.Login(ws)
	return nil
}

