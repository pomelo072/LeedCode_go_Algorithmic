package First_Bad_Version

import "sort"

var isBadVersion func(v int) bool

// FirstBadVersion1 官方解法
func FirstBadVersion1(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}

func FirstBadVersion2(n int) int {
	head := 1
	var half int
	for head <= n {
		half = head + (n-head)/2
		if isBadVersion(half) {
			n = half - 1
		} else {
			head = half + 1
		}
	}
	return head
}
