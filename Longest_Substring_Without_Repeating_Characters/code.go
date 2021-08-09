package Longest_Substring_Without_Repeating_Characters

func LengthOfLongestSubstring(s string) int {
	elementIndexMap := make(map[byte]int, 0)
	var bginIndex int
	var sum int
	for i := 0; i < len(s); i++ {
		if v, ok := elementIndexMap[s[i]]; ok {
			if v+1 > bginIndex {
				bginIndex = v + 1
			}
		}
		elementIndexMap[s[i]] = i
		subsutm := i - bginIndex + 1
		if subsutm > sum {
			sum = subsutm
		}
	}
	return sum
}
