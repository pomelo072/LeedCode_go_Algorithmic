package Two_Sum_Input_array_is_sorted

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums1 := []int{2, 7, 11, 15}
	t1 := 9
	nums2 := []int{2, 3, 4}
	t2 := 6
	nums3 := []int{-1, 0}
	t3 := -1
	v1 := TwoSum(nums1, t1)
	v2 := TwoSum(nums2, t2)
	v3 := TwoSum(nums3, t3)
	if assert.Equal(t, v1, []int{1, 2}) {
		t.Log(v1)
	} else {
		t.Log(v1)
		t.Error("v1 Error.")
	}
	if assert.Equal(t, v2, []int{1, 3}) {
		t.Log(v2)
	} else {
		t.Log(v2)
		t.Error("v2 Error.")
	}
	if assert.Equal(t, v3, []int{1, 2}) {
		t.Log(v3)
	} else {
		t.Log(v3)
		t.Error("v3 Error.")
	}
}
