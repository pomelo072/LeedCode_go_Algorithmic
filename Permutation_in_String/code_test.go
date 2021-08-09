package Permutation_in_String

import "testing"

func TestCheckInclusion(t *testing.T) {
	s1, s2, s3 := "ab", "eidbaooo", "eidboaoo"
	v1, v2 := CheckInclusion(s1, s2), CheckInclusion(s1, s3)
	if v1 && (!v2) {
		t.Log(v1)
		t.Log(v2)
	} else {
		t.Log(v1)
		t.Log(v2)
		t.Error("Error")
	}
}
