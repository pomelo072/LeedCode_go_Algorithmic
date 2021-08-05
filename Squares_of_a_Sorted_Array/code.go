package Squares_of_a_Sorted_Array

func SortedSquares(nums []int) []int {
	n := len(nums)
	newNums := make([]int, n, n)
	pointHead, pointTail := 0, n-1
	for i := n - 1; i >= 0; i-- {
		if h, t := nums[pointHead]*nums[pointHead], nums[pointTail]*nums[pointTail]; h > t {
			newNums[i] = h
			pointHead++
		} else {
			newNums[i] = t
			pointTail--
		}
	}
	return newNums
}
