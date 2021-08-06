package Move_Zeroes

func MoveZeroes(nums []int) {
	n := len(nums)
	numPoint, zeroPoint := -1, 0
	for i := 0; i < n; i++ {
		if nums[zeroPoint] != 0 {
			nums[numPoint+1] = nums[zeroPoint]
			numPoint++
		}
		zeroPoint++
	}
	for i := numPoint + 1; i < n; i++ {
		nums[i] = 0
	}
}
