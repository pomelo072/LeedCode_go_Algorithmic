package Binary_Search

func Search(nums []int, target int) int {
	head := 0
	tail := len(nums) - 1
	half := 0
	for head <= tail {
		half = head + (tail-head)/2
		if target == nums[half] {
			return half
		}
		if target < nums[half] {
			tail = half - 1
		} else {
			head = half + 1
		}
	}
	return -1
}
