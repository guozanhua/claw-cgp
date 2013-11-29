// Author: sheppard(ysf1026@gmail.com) 2013-11-29

package player

import (
	"fmt"
	pb "code.google.com/p/goprotobuf/proto"
	. "github.com/yangsf5/claw-cgp/proto"
)

var (
	handlers map[string] func(*Player)
)

func init() {
	handlers = make(map[string] func(*Player))
	RegisterHandler("*proto.CLPing", func(p *Player) {
		p.Send(&LCRetPong{Time: pb.Int32(222)})
	})
}

func RegisterHandler(proto string, handler func(*Player)) {
	_, ok := handlers[proto]
	if ok {
		panic(fmt.Sprintf("Register error, proto exist, proto=%s", proto))
	}

	handlers[proto] = handler
}
