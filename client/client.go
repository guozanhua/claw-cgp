// Author: sheppard(ysf1026@gmail.com) 2013-11-26

package main

import (
	"fmt"
	"time"
	"github.com/yangsf5/claw-cgp/client/player"
)

func main() {
	fmt.Println("Client starts ...")

	player := &player.Player{}
	player.Login()

	for {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Client stops ...")
}
