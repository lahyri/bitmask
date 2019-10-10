package bitmask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	a := uint16(0b1010101)
	a = Reset(a)
	assert.Equal(t, a, uint16(0))

}

func TestAdd(t *testing.T) {
	a := uint16(0)
	b := uint16(25)
	a = Add(a, b)
	assert.Equal(t, a, b)
}

func TestRemove(t *testing.T) {
	a := uint16(4)
	b := uint16(5)
	a = Remove(a, b)
	assert.Equal(t, a, uint16(0))
}
