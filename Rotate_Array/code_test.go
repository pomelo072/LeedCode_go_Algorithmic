package Rotate_Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	n1 := 3
	nums2 := []int{-1, -100, 3, 99}
	n2 := 2
	Rotate(nums1, n1)
	Rotate(nums2, n2)
	if assert.Equal(t, nums1, []int{5, 6, 7, 1, 2, 3, 4}) {
		t.Log(nums1)
	} else {
		t.Log(nums1)
		t.Error("v1 Error.")
	}
	if assert.Equal(t, nums2, []int{3, 99, -1, -100}) {
		t.Log(nums2)
	} else {
		t.Log(nums2)
		t.Error("v2 Error.")
	}
}
