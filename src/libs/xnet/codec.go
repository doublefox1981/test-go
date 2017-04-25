package xnet

import (
	"io"
)

// Packet TODO
type Packet struct {
	Len uint32
	Cmd uint32
	Seq uint32
	Msg interface{}
}

// Codec TODO
type Codec interface {
	Receive() (*Packet, error)
	Send(interface{}, uint32) error
	Close() error
}

// Protocol TODO
type Protocol interface {
	NewCodec(rw io.ReadWriteCloser) Codec
	GetPacket() *Packet
	PutPacket(p *Packet)
}
