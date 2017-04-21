package xnet

import (
	"sync"
)

const bucketNum = 32

// ClientMgr TODO
type ClientMgr struct {
	pods [bucketNum]clientPod
}

type clientPod struct {
	sync.RWMutex
	clients map[uint64]*Client
}

// NewClientMgr TODO
func NewClientMgr() *ClientMgr {
	m := &ClientMgr{}
	for i := 0; i < len(m.pods); i++ {
		m.pods[i].clients = make(map[uint64]*Client)
	}
	return m
}

// CreateClient TODO
func (m *ClientMgr) CreateClient(codec Codec) *Client {
	c := newClient(m, codec)
	pod := &m.pods[c.id%bucketNum]
	pod.Lock()
	defer pod.Unlock()
	pod.clients[c.id] = c
	return c
}

// GetClient TODO
func (m *ClientMgr) GetClient(id uint64) *Client {
	pod := &m.pods[id%bucketNum]
	pod.RLock()
	defer pod.RUnlock()
	c, _ := pod.clients[id]
	return c
}

// DeleteClient TODO
func (m *ClientMgr) DeleteClient(c *Client) {
	pod := &m.pods[c.id%bucketNum]
	pod.Lock()
	defer pod.Unlock()
	delete(pod.clients, c.id)
}
