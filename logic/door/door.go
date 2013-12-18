// Author: sheppard(ysf1026@gmail.com) 2013-12-13

package door

import (
	"net"
	"github.com/yangsf5/claw-cgp/logic/player"
)

func init() {
}

func Login(conn net.Conn) {
	client := &player.Player {
		Conn: conn,
	}
	client.Start()
}

