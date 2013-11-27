// Author: sheppard(ysf1026@gmail.com) 2013-11-27

package player

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	pb "code.google.com/p/goprotobuf/proto"
	myNet "github.com/yangsf5/claw-cgp/logic/net"
	. "github.com/yangsf5/claw-cgp/proto"
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
	_, err := net.Dial("tcp", "127.0.0.1:1104")
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
		//pbMsg, err := myNet.Decode(msg)
		myNet.Decode(msg)
	}
	defer p.conn.Close()
}

func (p *Player) Send(msg pb.Message) {
	msgBuf, err := myNet.Encode(msg)
	p.checkNetError(err)
	_, err = p.conn.Write(msgBuf)
	p.checkNetError(err)
}

func (p *Player) Recv() []byte {
	// Read total length
	var totalLen uint16
	var totalLenBuf [2]byte
	_, err := p.conn.Read(totalLenBuf[:])
	p.checkNetError(err)
	var totalLenBuffer *bytes.Buffer = bytes.NewBuffer(totalLenBuf[:])
	binary.Read(totalLenBuffer, binary.BigEndian, &totalLen)

	// Read the rest of message
	restBuf := make([]byte, totalLen)
	_, err = p.conn.Read(restBuf[:])
	p.checkNetError(err)

	msg := new(bytes.Buffer)
	msg.Write(totalLenBuf[:])
	msg.Write(restBuf[:])

	return msg.Bytes()
}

func (p *Player) checkNetError(err error) {
	if err != nil {
		fmt.Println("Net error:", err.Error())
		p.Disconnected = true
		// TODO check
		panic(err)
	}
}
