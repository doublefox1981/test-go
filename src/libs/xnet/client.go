package xnet

import (
	"sync/atomic"
)

var allocClientID uint64

// Client TODO
type Client struct {
	id    uint64
	codec Codec
	mgr   *ClientMgr
}

func newClient(mgr *ClientMgr, codec Codec) *Client {
	c := &Client{
		codec: codec,
		mgr:   mgr,
		id:    atomic.AddUint64(&allocClientID, 1),
	}

	return c
}

// Receive TODO
func (c *Client) Receive() (*Packet, error) {
	return c.codec.Receive()
}

// Close TODO
func (c *Client) Close() error {
	c.codec.Close()
	c.mgr.DeleteClient(c)
	return nil
}
