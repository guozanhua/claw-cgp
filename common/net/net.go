// Author: sheppard(ysf1026@gmail.com) 2013-11-22

package net

import (
	"bytes"
	"encoding/binary"
	"net"
)

func Handle() {
}

func Recv(conn net.Conn) ([]byte, error) {
	// Read total length
	var totalLen uint16
	var totalLenBuf [2]byte
	_, err := conn.Read(totalLenBuf[:])
	if err != nil {
		return nil, err
	}
	var totalLenBuffer *bytes.Buffer = bytes.NewBuffer(totalLenBuf[:])
	binary.Read(totalLenBuffer, binary.BigEndian, &totalLen)

	// Read the rest of message
	restBuf := make([]byte, totalLen)
	_, err = conn.Read(restBuf[:])
	if err != nil {
		return nil, err
	}

	msg := new(bytes.Buffer)
	msg.Write(totalLenBuf[:])
	msg.Write(restBuf[:])

	return msg.Bytes(), nil
}
