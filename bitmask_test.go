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
	b := uint16(0b11010111)
	a = Add(a, b)
	assert.Equal(t, a, b)
}

func TestRemove(t *testing.T) {
	a := uint16(0b100)
	b := uint16(0b101)
	a = Remove(a, b)
	assert.Equal(t, a, uint16(0))
}

func TestHas(t *testing.T) {
	a := uint16(0b1011100)
	b := uint16(0b1000000)
	c := uint16(0b1000010)
	assert.True(t, Has(a, b))
	assert.False(t, Has(a, c))
}

