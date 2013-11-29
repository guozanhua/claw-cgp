// Author: sheppard(ysf1026@gmail.com) 2013-11-22

package player

import (
	"fmt"
	"net"
	pb "code.google.com/p/goprotobuf/proto"
	myNet "github.com/yangsf5/claw-cgp/common/net"
)

type Player struct {
	Conn net.Conn
	Disconnected bool
}

func (p *Player) Start() {
	fmt.Println("Player start")
	p.Disconnected = false
	go p.Tick()
}

func (p *Player) Stop() {
	p.Disconnected = true
}

func (p *Player) Tick() {
	for ; !p.Disconnected; {
		msgName, msg, err := myNet.Recv(p.Conn)
		p.checkNetError(err)
		if err != nil {
			break
		}
		fmt.Println("Recv msg", msgName)
		exec, ok := handlers[msgName]
		if ok {
			exec(p, msg)
		} else {
			fmt.Println("No handler")
		}
	}
	defer p.Conn.Close()
	fmt.Println("Client disconneted")
}

func (p *Player) Send(msg pb.Message) {
	err := myNet.Send(p.Conn, msg)
	p.checkNetError(err)
}

func (p *Player) checkNetError(err error) {
	if err != nil {
		fmt.Println("Net error:", err.Error())
		p.Disconnected = true
	}
}
