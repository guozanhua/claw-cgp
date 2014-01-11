// Author: sheppard(ysf1026@gmail.com) 2013-11-07

package main

import (
	"fmt"
	"net"
	"github.com/yangsf5/claw-cgp/app/logic/player"
	"github.com/yangsf5/claw-cgp/util"
)

func main() {
	fmt.Println("Logic starts ...")

	service := ":1104"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	util.CheckFatal(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	util.CheckFatal(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		client := &player.Player{
			Conn: conn,
		}
		go client.Start()
	}


	fmt.Println("Logic stops ...")
}
