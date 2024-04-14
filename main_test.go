package main_test

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestLearnSlices(t *testing.T) {
	slice := []int{2, 3, 4, 5, 11, 17}
	assert.Equal(t, len(slice), 6, "length should be 6")
	assert.Equal(t, cap(slice), 6, "capacity should be 6")

	slice = slice[:0]
	assert.Equal(t, len(slice), 0, "length should be 0")
	assert.Equal(t, cap(slice), 6, "capacity should be 6")

	slice = slice[:4]
	assert.Equal(t, len(slice), 4, "length should be 4")
	assert.Equal(t, cap(slice), 6, "capacity should be 6")

	slice = slice[2:]
	assert.Equal(t, len(slice), 2, "length should be 2")
	assert.Equal(t, cap(slice), 4, "capacity should be 4")
}
