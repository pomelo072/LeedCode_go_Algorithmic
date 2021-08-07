package Reverse_Words_in_a_String_III

import "strings"

func ReverseWords(s string) string {
	var b []byte
	length := len(s)
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for w := start; w < i; w++ {
			b = append(b, s[start+i-1-w])
		}
		for i < length && s[i] == ' ' {
			i++
			b = append(b, ' ')
		}
	}
	return string(b)
}

func ReverseWords2(s string) string {
	var b strings.Builder
	length := len(s)
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for w := start; w < i; w++ {
			b.WriteByte(s[start+i-1-w])
		}
		for i < length && s[i] == ' ' {
			i++
			b.WriteByte(' ')
		}
	}
	return b.String()
}
