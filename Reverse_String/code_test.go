package Reverse_String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	b1 := []byte{'h', 'e', 'l', 'l', 'o'}
	b2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	ReverseString(b1)
	ReverseString(b2)
	if assert.Equal(t, b1, []byte{'o', 'l', 'l', 'e', 'h'}) {
		t.Log(b1)
	} else {
		t.Log(b1)
		t.Error("b1 Error.")
	}
	if assert.Equal(t, b2, []byte{'h', 'a', 'n', 'n', 'a', 'H'}) {
		t.Log(b2)
	} else {
		t.Log(b2)
		t.Error("b2 Error.")
	}
}
