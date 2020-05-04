package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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
