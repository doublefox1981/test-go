package xnet

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"sync"

	"github.com/golang/protobuf/proto"
)

const (
	headLen = 8
)

var (
	// ErrPacketSize TODO
	ErrPacketSize = errors.New("pbcodec: packet size")
)

// Protobuf TODO
func Protobuf(n int) *PBProtocol {
	return &PBProtocol{
		maxPacketLen: n,
		pbPool:       make(map[uint16]*sync.Pool),
	}
}

// PBProtocol TODO
type PBProtocol struct {
	maxPacketLen int
	pbPool       map[uint16]*sync.Pool
}

// NewCodec TODO
func (p *PBProtocol) NewCodec(rw io.ReadWriteCloser) (Codec, error) {
	return &PBCodec{
		protocol: p,
		conn:     rw,
		br:       bufio.NewReaderSize(rw, p.maxPacketLen+headLen),
	}, nil
}

// RegisterPool TODO
func (p *PBProtocol) RegisterPool(cmd uint16, msg proto.Message, f func() interface{}) {
	p.pbPool[cmd] = &sync.Pool{
		New: f,
	}
}

// PBCodec TODO
type PBCodec struct {
	whead    [headLen]byte
	rhead    []byte
	body     []byte
	conn     io.ReadWriteCloser
	br       *bufio.Reader
	protocol *PBProtocol
}

func (c *PBCodec) readFull(n int) ([]byte, error) {
	for {
		d, err := c.br.Peek(n)
		if len(d) >= n || err != nil {
			c.br.Discard(len(d))
			return d, err
		}
	}
}

// Receive TODO
func (c *PBCodec) Receive() (*Packet, error) {
	var err error
	if c.rhead, err = c.readFull(headLen); err != nil {
		return nil, err
	}
	pack := &Packet{}
	pack.len = binary.BigEndian.Uint16(c.rhead)
	if int(pack.len) > c.protocol.maxPacketLen {
		return nil, ErrPacketSize
	}
	pack.cmd = binary.BigEndian.Uint16(c.rhead[2:])
	pack.seq = binary.BigEndian.Uint32(c.rhead[4:])
	c.body, err = c.readFull(int(pack.len) - headLen)
	return pack, err
}

// Send TODO
func (c *PBCodec) Send(interface{}) error {
	return nil
}

// Close TODO
func (c *PBCodec) Close() error {
	return c.conn.Close()
}
