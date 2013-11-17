// Author: sheppard(ysf1026@gmail.com) 2013-11-16

package net

import (
	"fmt"
	"testing"
	pb "code.google.com/p/goprotobuf/proto"
	. "github.com/yangsf5/claw-cgp/proto"
)

func TestEncode(t *testing.T) {
	ping := &CLPing{
		Time: pb.Int32(12),
	}

	data, err := Encode(ping)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))
}

func TestDecode(t *testing.T) {
}
