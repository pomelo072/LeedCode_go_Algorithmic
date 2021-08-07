package Reverse_Words_in_a_String_III

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	v := ReverseWords(s)
	if assert.Equal(t, v, "s'teL ekat edoCteeL tsetnoc") {
		t.Log(v)
	} else {
		t.Log(v)
		t.Error("v Error.")
	}
}

func BenchmarkReverseWords(b *testing.B) {
	s := "Let's take LeetCode contest"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ReverseWords(s)
	}
	b.StopTimer()
}

func BenchmarkReverseWords2(b *testing.B) {
	s := "Let's take LeetCode contest"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ReverseWords2(s)
	}
	b.StopTimer()
}
