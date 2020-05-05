package server

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockHandler struct{}

func (m *mockHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "hello")
}

func TestServer_ServeHTTP(t *testing.T) {
	s := &Server{
		handler: &mockHandler{},
	}
	r := httptest.NewRecorder()
	s.ServeHTTP(r, nil)

	want := "hello"
	got := r.Body.String()
	assert.Equal(t, want, got)
}

func TestServer_IsAlive(t *testing.T) {
	ln, err := net.Listen("tcp4", ":1234")
	assert.NoError(t, err)

	u, _ := url.Parse("http://0.0.0.0:1234")
	server := &Server{URL: u}
	assert.True(t, server.IsAlive())

	ln.Close()
	assert.False(t, server.IsAlive())
}
