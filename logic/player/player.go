// Author: sheppard(ysf1026@gmail.com) 2013-11-22

package player

import (
	"fmt"
	"net"
)

type Player struct {
	Conn net.Conn
}

func (p *Player) Start() {
	fmt.Println("Player start")
}
