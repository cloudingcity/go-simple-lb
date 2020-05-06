package server

import (
	"container/list"
	"fmt"
	"net/http"
	"sync"
)

type Controller struct {
	servers map[int]*Server
	upIDs   *list.List
	downIDs *list.List
	mux     sync.Mutex
}

func NewController() *Controller {
	return &Controller{
		upIDs:   list.New(),
		downIDs: list.New(),
	}
}

func (c *Controller) SetServers(servers []*Server) {
	c.servers = make(map[int]*Server, len(servers))

	for i, server := range servers {
		id := i + 1
		c.servers[id] = server
		c.upIDs.PushBack(id)
	}
}

func (c *Controller) GetNext() *Server {
	id := c.getNextID()
	if id == 0 {
		return nil
	}
	return c.servers[id]
}

func (c *Controller) getNextID() int {
	defer c.mux.Unlock()
	c.mux.Lock()
	if e := c.upIDs.Front(); e != nil {
		c.upIDs.MoveToBack(e)
		return e.Value.(int)
	}
	return 0
}

func (c *Controller) HTTPHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if server := c.GetNext(); server != nil {
			server.ServeHTTP(rw, req)
			return
		}
		http.Error(rw, "Service unavailable", http.StatusServiceUnavailable)
	})
}

func (c *Controller) HealthCheck() []string {
	defer c.mux.Unlock()
	c.mux.Lock()

	var (
		msgs []string
		next *list.Element
	)
	for e := c.upIDs.Front(); e != nil; e = next {
		next = e.Next()
		id := e.Value.(int)
		server := c.servers[id]
		if server.IsAlive() {
			continue
		}
		c.downIDs.PushBack(id)
		c.upIDs.Remove(e)
		msgs = append(msgs, fmt.Sprintf("[%s] down", server.url))
	}
	for e := c.downIDs.Front(); e != nil; e = next {
		next = e.Next()
		id := e.Value.(int)
		server := c.servers[id]
		if !server.IsAlive() {
			continue
		}
		c.upIDs.PushBack(id)
		c.downIDs.Remove(e)
		msgs = append(msgs, fmt.Sprintf("[%s] up", server.url))
	}
	return msgs
}
