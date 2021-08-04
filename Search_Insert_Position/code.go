package Search_Insert_Position

func SearchInsert(nums []int, target int) int {
	head := 0
	tail := len(nums) - 1
	var half int
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
	return head
}
