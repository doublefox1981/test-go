package xnet

import (
	"net"
	"sync"
)

const bucketNum = 32

// SessionMgr TODO
type SessionMgr struct {
	pods [bucketNum]sessionPod
	svr  *Server
}

type sessionPod struct {
	sync.RWMutex
	sessions map[uint64]*Session
}

// NewSessionMgr TODO
func NewSessionMgr(svr *Server) *SessionMgr {
	m := &SessionMgr{svr: svr}
	for i := 0; i < len(m.pods); i++ {
		m.pods[i].sessions = make(map[uint64]*Session)
	}
	return m
}

// CreateSession TODO
func (m *SessionMgr) CreateSession(conn net.Conn, codec Codec, sndChanSize int) *Session {
	c := newSession(m, conn, codec, sndChanSize)
	pod := &m.pods[c.id%bucketNum]
	pod.Lock()
	defer pod.Unlock()
	pod.sessions[c.id] = c
	return c
}

// GetSession TODO
func (m *SessionMgr) GetSession(id uint64) *Session {
	pod := &m.pods[id%bucketNum]
	pod.RLock()
	defer pod.RUnlock()
	c, _ := pod.sessions[id]
	return c
}

// DeleteSession TODO
func (m *SessionMgr) DeleteSession(c *Session) {
	pod := &m.pods[c.id%bucketNum]
	pod.Lock()
	defer pod.Unlock()
	delete(pod.sessions, c.id)
}
