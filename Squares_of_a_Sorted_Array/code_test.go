package Squares_of_a_Sorted_Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	nums1 := []int{-4, -1, 0, 3, 10}
	nums2 := []int{-7, -3, 2, 3, 11}
	v1 := SortedSquares(nums1)
	v2 := SortedSquares(nums2)
	if assert.Equal(t, v1, []int{0, 1, 9, 16, 100}) {
		t.Log(v1)
	} else {
		t.Error("v1 Error.")
	}
	if assert.Equal(t, v2, []int{4, 9, 9, 49, 121}) {
		t.Log(v2)
	} else {
		t.Error("v2 Error.")
	}
}
func BenchmarkSortedSquares(b *testing.B) {
	nums1 := []int{-4, -1, 0, 3, 10}
	nums2 := []int{-7, -3, 2, 3, 11}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SortedSquares(nums1)
		_ = SortedSquares(nums2)
	}
	b.StopTimer()
}
