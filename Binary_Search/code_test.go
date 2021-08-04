package Binary_Search

import "testing"

func TestSearch(t *testing.T) {
	l := []int{-1, 0, 3, 5, 9, 12}
	t1, t2 := 9, 2
	v1 := Search(l, t1)
	v2 := Search(l, t2)
	t.Logf("v1: %d Actully: %d", v1, 4)
	t.Logf("v2: %d Actully: %d", v2, -1)
}
