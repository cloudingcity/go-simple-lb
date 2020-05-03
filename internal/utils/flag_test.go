package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagURL_Set(t *testing.T) {
	f := &FlagURL{}
	err := f.Set("http://foo,http://bar,http://baz")
	assert.NoError(t, err)

	want := "[http://foo http://bar http://baz]"
	got := f.String()
	assert.Equal(t, want, got)
}

func TestFlagURL_Set_Error(t *testing.T) {
	f := &FlagURL{}
	err := f.Set("foo,bar,baz")
	assert.Error(t, err)
}
