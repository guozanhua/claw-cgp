// Author: sheppard(ysf1026@gmail.com) 2013-11-27

package player

import (
	"fmt"
	"net"
	pb "code.google.com/p/goprotobuf/proto"
	myNet "github.com/yangsf5/claw-cgp/common/net"
	. "github.com/yangsf5/claw-cgp/common/proto"
	"github.com/yangsf5/claw-cgp/util"
)

type Player struct {
	conn net.Conn
	Disconnected bool
}

func (p *Player) Start() {
	p.Disconnected = false
	p.Login()
	go p.Tick()
}

func (p *Player) Stop() {
	p.Disconnected = true
}

func (p *Player) Login() {
	conn, err := net.Dial("tcp", "127.0.0.1:1104")
	p.conn = conn
	util.CheckFatal(err)
	fmt.Println("Connnected")
}

func (p *Player) Ping() {
	ping := CLPing{
		Time: pb.Int32(555),
	}
	p.Send(&ping)
}

func (p *Player) Tick() {
	for ; !p.Disconnected; {
		msg := p.Recv()
		msgName, _, err := myNet.Decode(msg)
		if err != nil {
			fmt.Println("Message decode error", err)
		}
		fmt.Println("Recv msg", msgName)
		exec, ok := handlers[msgName]
		if ok {
			exec(p)
		} else {
			fmt.Println("No handler")
		}
	}
	defer p.conn.Close()
}

func (p *Player) Send(msg pb.Message) {
	err := myNet.Send(p.conn, msg)
	p.checkNetError(err)
}

func (p *Player) Recv() []byte {
	buf, err := myNet.Recv(p.conn)
	p.checkNetError(err)
	return buf
}

func (p *Player) checkNetError(err error) {
	if err != nil {
		fmt.Println("Net error:", err.Error())
		p.Disconnected = true
		// TODO check
		panic(err)
	}
}
