package Two_Sum_Input_array_is_sorted

// 0ms解法
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))
	for i, v := range numbers {
		another := target - v
		if m[another] > 0 {
			return []int{m[another], i + 1}
		}
		m[v] = i + 1
	}
	return nil
}

func TwoSum(numbers []int, target int) []int {
	search := func(nums []int, tg int) int {
		head := 0
		tail := len(nums) - 1
		half := 0
		for head <= tail {
			half = head + (tail-head)/2
			if tg == nums[half] {
				return half + 1
			}
			if tg < nums[half] {
				tail = half - 1
			} else {
				head = half + 1
			}
		}
		return -1
	}
	for i, v := range numbers {
		tgNum := target - v
		tgIndex := search(numbers[i+1:], tgNum)
		if tgIndex == -1 {
			continue
		} else {
			result := []int{i + 1, tgIndex + i + 1}
			return result
		}
	}
	return nil
}
