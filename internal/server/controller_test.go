package server

import (
	"net"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController_GetNext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		u1, _ := url.Parse("http://example1.com")
		u2, _ := url.Parse("http://example2.com")
		u3, _ := url.Parse("http://example3.com")
		s1 := &Server{url: u1}
		s2 := &Server{url: u2}
		s3 := &Server{url: u3}

		c := NewController()
		c.SetServers([]*Server{s1, s2, s3})

		assert.Equal(t, s1, c.GetNext())
		assert.Equal(t, s2, c.GetNext())
		assert.Equal(t, s3, c.GetNext())
	})
	t.Run("failed", func(t *testing.T) {
		c := NewController()
		assert.Nil(t, c.GetNext())
	})
}

func TestController_HealthCheck(t *testing.T) {
	u1, _ := url.Parse("http://localhost:1234")
	u2, _ := url.Parse("http://localhost:1235")
	u3, _ := url.Parse("http://localhost:1236")
	s1 := &Server{url: u1}
	s2 := &Server{url: u2}
	s3 := &Server{url: u3}

	c := NewController()
	c.SetServers([]*Server{s1, s2, s3})

	c.HealthCheck()
	assert.Equal(t, 0, c.upIDs.Len())
	assert.Equal(t, 3, c.downIDs.Len())

	ln, err := net.Listen("tcp4", "localhost:1234")
	assert.NoError(t, err)

	c.HealthCheck()
	assert.Equal(t, 1, c.upIDs.Len())
	assert.Equal(t, 2, c.downIDs.Len())

	ln.Close()
}
