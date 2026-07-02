package http2

import "sync"

// ClientConn represents an HTTP/2 connection.
type ClientConn struct {
	mu     sync.Mutex
	closed bool
}

// HandleGoAway processes the GOAWAY frame.
func (cc *ClientConn) HandleGoAway(pool *ClientConnPool, key string) {
	cc.mu.Lock()
	cc.closed = true
	cc.mu.Unlock()
	pool.RemoveConn(key, cc)
}

// ClientConnPool manages active connections.
type ClientConnPool struct {
	mu    sync.Mutex
	conns map[string]*ClientConn
}

// RemoveConn removes a connection from the pool.
func (p *ClientConnPool) RemoveConn(key string, cc *ClientConn) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.conns[key] == cc {
		delete(p.conns, key)
	}
}