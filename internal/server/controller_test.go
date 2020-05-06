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

		c := NewController()
		c.SetupServers(u1, u2, u3)

		assert.Equal(t, u1, c.GetNext().url)
		assert.Equal(t, u2, c.GetNext().url)
		assert.Equal(t, u3, c.GetNext().url)
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

	c := NewController()
	c.SetupServers(u1, u2, u3)

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
