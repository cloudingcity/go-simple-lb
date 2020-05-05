package server

import (
	"net"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPool_GetNext(t *testing.T) {
	t.Run("polling success", func(t *testing.T) {
		u1, _ := url.Parse("http://example1.com")
		u2, _ := url.Parse("http://example2.com")
		u3, _ := url.Parse("http://example3.com")
		s1 := &Server{url: u1}
		s2 := &Server{url: u2}
		s3 := &Server{url: u3}

		pool := NewPool()
		pool.Put(s1)
		pool.Put(s2)
		pool.Put(s3)

		assert.Equal(t, s1, pool.GetNext())
		assert.Equal(t, s2, pool.GetNext())
		assert.Equal(t, s3, pool.GetNext())
	})
	t.Run("polling failed", func(t *testing.T) {
		pool := NewPool()
		assert.Nil(t, pool.GetNext())
	})
}

func TestPool_HealthCheck(t *testing.T) {
	u1, _ := url.Parse("http://0.0.0.0:1234")
	u2, _ := url.Parse("http://0.0.0.0:1235")
	u3, _ := url.Parse("http://0.0.0.0:1236")
	s1 := &Server{url: u1}
	s2 := &Server{url: u2}
	s3 := &Server{url: u3}

	pool := NewPool()
	pool.Put(s1)
	pool.Put(s2)
	pool.Put(s3)

	pool.HealthCheck()
	assert.Equal(t, 0, pool.servers.Len())
	assert.Equal(t, 3, pool.downServers.Len())

	ln, err := net.Listen("tcp4", ":1234")
	assert.NoError(t, err)

	pool.HealthCheck()
	assert.Equal(t, 1, pool.servers.Len())
	assert.Equal(t, 2, pool.downServers.Len())

	ln.Close()
}
