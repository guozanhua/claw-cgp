// Author: sheppard(ysf1026@gmail.com) 2013-11-16

package net

import (
	"fmt"
	"testing"
	pb "code.google.com/p/goprotobuf/proto"
	. "github.com/yangsf5/claw-cgp/common/proto"
)

func TestEncode(t *testing.T) {
	ping := &CLPing{
		Time: pb.Int32(12),
	}

	data, err := Encode(ping)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))

	var ping2 pb.Message
	ping2, err = Decode(data)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ping2)
}

func TestDecode(t *testing.T) {
}
