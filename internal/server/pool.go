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
	defer p.mux.Unlock()
	p.mux.Lock()

	if el := p.servers.Front(); el != nil {
		p.servers.MoveToBack(el)
		return el.Value.(*Server)
	}
	return nil
}
