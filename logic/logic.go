// Author: sheppard(ysf1026@gmail.com) 2013-11-07

package main

import (
	"fmt"
	pb "code.google.com/p/goprotobuf/proto"
	//"github.com/yangsf5/claw-cgp/logic/player"
	. "github.com/yangsf5/claw-cgp/proto"
)

func main() {
	fmt.Println("Logic starts ...")

	ping := &CLPing{
		Time: pb.Int32(12),
	}

	data, _ := pb.Marshal(ping)

	ping2 := new(CLPing)
	pb.Unmarshal(data, ping2)

	fmt.Println(ping2)


	fmt.Println("Logic stops ...")
}
