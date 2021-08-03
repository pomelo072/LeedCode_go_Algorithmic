package Longest_Common_Prefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs1 := []string{"flower", "flow", "flight"}
	strs2 := []string{"dog", "racecar", "car"}
	v1 := LongestCommonPrefix(strs1)
	t.Logf("v1 : %s Actully: %s", v1, "fl")
	v2 := LongestCommonPrefix(strs2)
	t.Logf("v1 : %s Actully: %s", v2, "")
}
