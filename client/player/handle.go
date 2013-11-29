// Author: sheppard(ysf1026@gmail.com) 2013-11-29

package player

import (
	"fmt"
)

var (
	handlers map[string] func(*Player)
)

func init() {
	handlers = make(map[string] func(*Player))
	RegisterHandler("LCRetPong", func(p *Player) {
		fmt.Println("haha, nothing")
	})
}

func RegisterHandler(proto string, handler func(*Player)) {
	_, ok := handlers[proto]
	if ok {
		panic(fmt.Sprintf("Register error, proto exist, proto=%s", proto))
	}

	handlers[proto] = handler
}
