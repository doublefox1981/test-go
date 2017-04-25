package xnet

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"sync"

	"reflect"

	"github.com/golang/protobuf/proto"
)

type pbPacketLenType uint16

// PBCmdType TODO
type PBCmdType uint16

const (
	headLen = 8
)

var (
	// ErrPacketSize TODO
	ErrPacketSize = errors.New("pbcodec: packet size")
	// ErrPacketCMD TODO
	ErrPacketCMD = errors.New("pbcodec: packet cmd")
	// ErrUnmarshal TODO
	ErrUnmarshal = errors.New("pbcodec: unmarshaler")
	// ErrUnregister TODO
	ErrUnregister = errors.New("pbcodec: unregister")
)

// Protobuf TODO
func Protobuf(n uint32) *PBProtocol {
	return &PBProtocol{
		maxPacketLen: n,
		pbPool:       make(map[PBCmdType]*sync.Pool),
		pbCmd:        make(map[reflect.Type]PBCmdType),
		packetPool: &sync.Pool{
			New: func() interface{} {
				return &Packet{}
			},
		},
	}
}

// PBProtocol TODO
type PBProtocol struct {
	maxPacketLen uint32
	pbPool       map[PBCmdType]*sync.Pool
	pbCmd        map[reflect.Type]PBCmdType
	packetPool   *sync.Pool
}

// NewCodec TODO
func (p *PBProtocol) NewCodec(rw io.ReadWriteCloser) Codec {
	wb := make([]byte, headLen+p.maxPacketLen)
	return &PBCodec{
		protocol: p,
		conn:     rw,
		br:       bufio.NewReaderSize(rw, int(p.maxPacketLen+headLen)),
		wbuf:     wb,
		pbbuf:    proto.NewBuffer(wb[headLen:]),
	}
}

// GetPacket TODO
func (p *PBProtocol) GetPacket() *Packet {
	return p.packetPool.Get().(*Packet)
}

// PutPacket TODO
func (p *PBProtocol) PutPacket(pack *Packet) {
	p.packetPool.Put(pack)
}

type pBMsgCtor func() interface{}

// RegisterPB TODO
func (p *PBProtocol) RegisterPB(cmd PBCmdType, pbm proto.Message, ctor pBMsgCtor) {
	p.pbPool[cmd] = &sync.Pool{
		New: ctor,
	}
	p.pbCmd[reflect.TypeOf(pbm)] = cmd
}

// CreatePB TODO
func (p *PBProtocol) CreatePB(cmd PBCmdType) interface{} {
	v, ok := p.pbPool[cmd]
	if ok {
		return v.Get()
	}
	return nil
}

// PBCodec TODO
type PBCodec struct {
	wbuf     []byte
	rhead    []byte
	body     []byte
	seq      uint32
	conn     io.ReadWriteCloser
	br       *bufio.Reader
	protocol *PBProtocol
	pbbuf    *proto.Buffer
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
	pack := c.protocol.GetPacket()
	pack.Len = uint32(binary.BigEndian.Uint16(c.rhead))
	if pack.Len > c.protocol.maxPacketLen {
		return nil, ErrPacketSize
	}
	pack.Cmd = uint32(binary.BigEndian.Uint16(c.rhead[2:]))
	p := c.protocol.CreatePB(PBCmdType(pack.Cmd))
	if p == nil {
		return nil, ErrPacketCMD
	}
	pm := p.(proto.Message)
	pack.Seq = binary.BigEndian.Uint32(c.rhead[4:])
	c.body, err = c.readFull(int(pack.Len) - headLen)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(c.body, pm)
	if err != nil {
		return nil, err
	}
	pack.Msg = pm
	return pack, nil
}

// Send TODO
func (c *PBCodec) Send(p interface{}, seq uint32) error {
	type marshalToer interface {
		MarshalTo(dAtA []byte) (int, error)
	}
	var (
		ok  bool
		m   marshalToer
		cmd PBCmdType
	)
	cmd, ok = c.protocol.pbCmd[reflect.TypeOf(p)]
	if !ok {
		return ErrUnregister
	}
	m, ok = p.(marshalToer)
	if !ok {
		return ErrUnmarshal
	}
	i, err := m.MarshalTo(c.wbuf[headLen:])
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint16(c.wbuf[0:2], uint16(headLen+i))
	binary.BigEndian.PutUint16(c.wbuf[2:4], uint16(cmd))
	if seq != 0 {
		binary.BigEndian.PutUint32(c.wbuf[4:8], seq)
	} else {
		c.seq++
		binary.BigEndian.PutUint32(c.wbuf[4:8], c.seq)
	}

	_, err = c.conn.Write(c.wbuf[:headLen+i])
	return err
}

// Close TODO
func (c *PBCodec) Close() error {
	return c.conn.Close()
}
