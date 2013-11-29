// Author: sheppard(ysf1026@gmail.com) 2013-11-22

package player

import (
	"bytes"
	"encoding/binary"
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
	// Read total length
	var totalLen uint16
	var totalLenBuf [2]byte
	_, err := p.Conn.Read(totalLenBuf[:])
	p.checkNetError(err)
	var totalLenBuffer *bytes.Buffer = bytes.NewBuffer(totalLenBuf[:])
	binary.Read(totalLenBuffer, binary.BigEndian, &totalLen)

	// Read the rest of message
	restBuf := make([]byte, totalLen)
	_, err = p.Conn.Read(restBuf[:])
	p.checkNetError(err)

	msg := new(bytes.Buffer)
	msg.Write(totalLenBuf[:])
	msg.Write(restBuf[:])
	fmt.Println(msg)

	return msg.Bytes()
}

func (p *Player) checkNetError(err error) {
	if err != nil {
		fmt.Println("Net error:", err.Error())
		p.Disconnected = true
	}
}
