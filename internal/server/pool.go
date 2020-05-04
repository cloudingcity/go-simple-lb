package server

import (
	"container/list"
	"sync"
)

type Pool struct {
	servers *list.List
	mux     sync.Mutex
}

func NewPool() *Pool {
	return &Pool{
		servers: list.New(),
	}
}

func (p *Pool) Put(server *Server) {
	p.mux.Lock()
	p.servers.PushBack(server)
	p.mux.Unlock()
}

func (p *Pool) GetNext() *Server {
	p.mux.Lock()
	el := p.servers.Front()
	p.servers.MoveToBack(el)
	p.mux.Unlock()
	return el.Value.(*Server)
}
