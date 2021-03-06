// Author: sheppard(ysf1026@gmail.com) 2013-11-29

package player

import (
	"fmt"
	pb "code.google.com/p/goprotobuf/proto"
	. "github.com/yangsf5/claw-cgp/app/common/proto"
)

var (
	handlers map[string] func(*Player, pb.Message)
)

func init() {
	handlers = make(map[string] func(*Player, pb.Message))
	RegisterHandler("CLPing", func(p *Player, msg pb.Message) {
		p.Send(&LCRetPong{Time: pb.Int32(222)})
	})
}

func RegisterHandler(proto string, handler func(*Player, pb.Message)) {
	_, ok := handlers[proto]
	if ok {
		panic(fmt.Sprintf("Register error, proto exist, proto=%s", proto))
	}

	handlers[proto] = handler
}
