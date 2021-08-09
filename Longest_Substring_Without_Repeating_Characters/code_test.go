package Longest_Substring_Without_Repeating_Characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s1, s2, s3, s4 := "abcabcbb", "bbbbb", "pwwkew", ""
	v1, v2, v3, v4 := LengthOfLongestSubstring(s1), LengthOfLongestSubstring(s2), LengthOfLongestSubstring(s3), LengthOfLongestSubstring(s4)
	if b1, b2, b3, b4 := assert.Equal(t, 3, v1), assert.Equal(t, 1, v2), assert.Equal(t, 3, v3), assert.Equal(t, 0, v4); b1 && b2 && b3 && b4 {
		t.Log(b1)
		t.Log(b2)
		t.Log(b3)
		t.Log(b4)
	} else {
		t.Log(b1)
		t.Log(b2)
		t.Log(b3)
		t.Log(b4)
		t.Error("Error")
	}
}
