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
		msg := p.Recv()
		msgName, pbMsg, err := myNet.Decode(msg)
		if err != nil {
			fmt.Println("Message decode error", err)
		}
		fmt.Println("Recv msg", msgName)
		exec, ok := handlers[msgName]
		if ok {
			exec(p, pbMsg)
		} else {
			fmt.Println("No handler")
		}
	}
	defer p.Conn.Close()
	fmt.Println("Client disconneted")
}

func (p *Player) Send(msg pb.Message) {
	msgBuf, err := myNet.Encode(msg)
	p.checkNetError(err)
	_, err = p.Conn.Write(msgBuf)
	p.checkNetError(err)
}

func (p *Player) Recv() []byte {
	buf, err := myNet.Recv(p.Conn)
	p.checkNetError(err)
	return buf
}

func (p *Player) checkNetError(err error) {
	if err != nil {
		fmt.Println("Net error:", err.Error())
		p.Disconnected = true
	}
}
