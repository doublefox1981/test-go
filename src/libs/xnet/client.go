package xnet

import "net"

// Connect TODO
func Connect(network, addr string, protocol Protocol, sndChanSize int) (*Session, error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return newSession(nil, conn, protocol.NewCodec(conn), sndChanSize), nil
}
