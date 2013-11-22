// Author: sheppard(ysf1026@gmail.com) 2013-11-16

package net

import (
	"bytes"
	"encoding/binary"
	"errors"
	"reflect"
	pb "code.google.com/p/goprotobuf/proto"
)

func Encode(msg pb.Message) ([]byte, error) {
	data, err := pb.Marshal(msg)
	if err != nil {
		return nil, err
	}
	msgName := reflect.TypeOf(msg).String()
	msgName = msgName[7:] //TODO better way to remove *proto.

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, byte(len(msgName)))
	if err != nil {
		return nil, err
	}
	buf.Write([]byte(msgName))

	err = binary.Write(buf, binary.BigEndian, uint16(len(data)))
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	return buf.Bytes(), nil
}

func Decode(input []byte) (pb.Message, error) {
	buf := bytes.NewBuffer(input)

	var nameLen byte
	err := binary.Read(buf, binary.BigEndian, &nameLen)
	if err != nil {
		return nil, err
	}
	msgName := string(buf.Next(int(nameLen)))

	var dataLen uint16
	err = binary.Read(buf, binary.BigEndian, &dataLen)
	if err != nil {
		return nil, err
	}
	data := buf.Next(int(dataLen))

	msgFunc, ok := protos[msgName]
	if ok {
		msg := msgFunc()
		pb.Unmarshal(data, msg)
		return msg, nil
	}
	return nil, errors.New("Message not registed")
}
