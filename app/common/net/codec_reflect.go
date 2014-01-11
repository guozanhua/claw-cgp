// Author: sheppard(ysf1026@gmail.com) 2013-11-19

package net

import (
	pb "code.google.com/p/goprotobuf/proto"
	. "github.com/yangsf5/claw-cgp/app/common/proto"
)

var (
	protos map[string] func() pb.Message
)

func init() {
	protos = make(map[string] func() pb.Message)
	protos["CLPing"] = func() pb.Message {
		return &CLPing{} //TODO check optional field default value?
	}
	protos["LCRetPong"] = func() pb.Message {
		return &LCRetPong{} //TODO check optional field default value?
	}
}
