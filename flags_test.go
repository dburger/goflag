package goflag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayFlags_Set(t *testing.T) {
	af := ArrayFlags{}
	af.Set("hello")
	af.Set("world")
	af.Set("bye")
	assert.Equal(t, ArrayFlags{"hello", "world", "bye"}, af)
}

func TestArrayFlags_String(t *testing.T) {
	af := ArrayFlags{"hello", "world"}
	assert.Equal(t, "hello,world", af.String())
}
