package Longest_Common_Prefix

func LongestCommonPrefix(strs []string) string {
	sPre := strs[0]
	for _, v := range strs {
		sPre = lmp(sPre, v)
		if len(sPre) == 0 {
			return ""
		}
	}
	return sPre
}

func lmp(a string, b string) string {
	length := min(len(a), len(b))
	s := 0
	for i := 0; i < length; i++ {
		if a[i] == b[i] {
			s++
		} else {
			break
		}
	}
	return a[:s]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
