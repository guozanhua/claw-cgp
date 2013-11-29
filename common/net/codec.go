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
	// Write total length = length of msgNameLen + msgNameLen + dataLen
	var totalLen uint16 = 1 + uint16(len(msgName)) + uint16(len(data))
	err = binary.Write(buf, binary.BigEndian, totalLen)
	if err != nil {
		return nil, err
	}

	// Write name length
	err = binary.Write(buf, binary.BigEndian, byte(len(msgName)))
	if err != nil {
		return nil, err
	}

	// Write name of message
	buf.Write([]byte(msgName))

	// Write data of message
	buf.Write(data)

	return buf.Bytes(), nil
}

func Decode(input []byte) (pb.Message, error) {
	buf := bytes.NewBuffer(input)

	// Read total length
	var totalLen uint16
	err := binary.Read(buf, binary.BigEndian, &totalLen)
	if err != nil {
		return nil, err
	}

	// Read name length
	var nameLen byte
	err = binary.Read(buf, binary.BigEndian, &nameLen)
	if err != nil {
		return nil, err
	}

	// Read name of message
	msgName := string(buf.Next(int(nameLen)))

	// Read data of message = totalLen - length of nameLen - nameLen
	dataLen := int(totalLen) - 1 - int(nameLen)
	data := buf.Next(dataLen)

	msgFunc, ok := protos[msgName]
	if ok {
		msg := msgFunc()
		pb.Unmarshal(data, msg)
		return msg, nil
	}
	return nil, errors.New("Message not registed")
}
