package server

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPool_GetNext(t *testing.T) {
	u1, _ := url.Parse("http://example1.com")
	u2, _ := url.Parse("http://example2.com")
	u3, _ := url.Parse("http://example3.com")
	s1 := &Server{URL: u1}
	s2 := &Server{URL: u2}
	s3 := &Server{URL: u3}

	pool := NewPool()
	pool.Put(s1)
	pool.Put(s2)
	pool.Put(s3)

	assert.Equal(t, s1, pool.GetNext())
	assert.Equal(t, s2, pool.GetNext())
	assert.Equal(t, s3, pool.GetNext())
}
