package server

import (
	"container/list"
	"fmt"
	"sync"
)

type Pool struct {
	servers     *list.List
	downServers *list.List
	mux         sync.Mutex
}

func NewPool() *Pool {
	return &Pool{
		servers:     list.New(),
		downServers: list.New(),
	}
}

func (p *Pool) Put(server *Server) {
	p.mux.Lock()
	p.servers.PushBack(server)
	p.mux.Unlock()
}

func (p *Pool) GetNext() *Server {
	defer p.mux.Unlock()
	p.mux.Lock()

	if e := p.servers.Front(); e != nil {
		p.servers.MoveToBack(e)
		return e.Value.(*Server)
	}
	return nil
}

func (p *Pool) HealthCheck() []string {
	defer p.mux.Unlock()
	p.mux.Lock()

	var (
		msgs []string
		next *list.Element
	)
	for e := p.servers.Front(); e != nil; e = next {
		next = e.Next()
		server := e.Value.(*Server)
		if server.IsAlive() {
			continue
		}
		p.downServers.PushBack(server)
		p.servers.Remove(e)
		msgs = append(msgs, fmt.Sprintf("[%s] down", server.serverURL))
	}
	for e := p.downServers.Front(); e != nil; e = next {
		next = e.Next()
		server := e.Value.(*Server)
		if !server.IsAlive() {
			continue
		}
		p.servers.PushBack(server)
		p.downServers.Remove(e)
		msgs = append(msgs, fmt.Sprintf("[%s] up", server.serverURL))
	}
	return msgs
}
