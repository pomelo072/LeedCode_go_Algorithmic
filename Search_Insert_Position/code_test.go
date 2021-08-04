package Search_Insert_Position

import "testing"

func TestSearchInsert(t *testing.T) {
	nums1 := []int{1, 3, 5, 6}
	nums2 := []int{1}
	t1, t2, t3, t4 := 5, 2, 7, 0
	v1 := SearchInsert(nums1, t1)
	v2 := SearchInsert(nums1, t2)
	v3 := SearchInsert(nums1, t3)
	v4 := SearchInsert(nums1, t4)
	v5 := SearchInsert(nums2, t4)
	t.Logf("%d %d", v1, 2)
	t.Logf("%d %d", v2, 1)
	t.Logf("%d %d", v3, 4)
	t.Logf("%d %d", v4, 0)
	t.Logf("%d %d", v5, 0)
}
