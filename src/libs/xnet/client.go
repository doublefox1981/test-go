package xnet

import "net"

// Connect TODO
func Connect(network, addr string, protocol Protocol, sndChanSize int) (*Client, error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &Client{
		codec: protocol.NewCodec(conn),
		conn:  conn,
	}, nil
}

// Client TODO
type Client struct {
	codec Codec
	conn  net.Conn
}

// Receive TODO
func (c *Client) Receive() (*Packet, error) {
	return c.codec.Receive()
}

// Send TODO
func (c *Client) Send(p interface{}) error {
	return c.codec.Send(p, 0)
}

// Close TODO
func (c *Client) Close() error {
	return c.codec.Close()
}
