package utils

import (
	"bytes"
	"encoding/binary"
	"io"
	"errors"
)

var Codec = NewCodec()

type codec struct {

}

func NewCodec() *codec {
	return &codec{}
}

func (c *codec) EncodePack(message []byte) ([]byte, error){

	var messageLen = int32(len(message))
	var pkg = new(bytes.Buffer)

	// write header
	err := binary.Write(pkg, binary.LittleEndian, messageLen)
	if err != nil {
		return nil, err
	}
	// write content
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

func (c *codec) DecodePack(read io.Reader) (string, error) {

	// read header
	var length int32
	err := binary.Read(read, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// read message
	var messageByte = make([]byte, length)
	n, err := read.Read(messageByte)
	if err != nil {
		return "", err
	}
	if int32(n) != length {
		return "", errors.New("decode message length error")
	}
	return string(messageByte), nil
}



