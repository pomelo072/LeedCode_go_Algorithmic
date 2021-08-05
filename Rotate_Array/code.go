package Rotate_Array

func Rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse := func(nums []int) {
		n := len(nums)
		for i := 0; i < n/2; i++ {
			nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		}
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
