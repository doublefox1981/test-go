package xnet

import (
	"io"
)

type Packet struct {
	len uint16
	cmd uint16
	seq uint32
	msg interface{}
}

type Codec interface {
	Receive() (*Packet, error)
	Send(interface{}) error
	Close() error
}

type Protocol interface {
	NewCodec(rw io.ReadWriteCloser) Codec
	GetPacket() *Packet
	PutPacket(p *Packet)
}
