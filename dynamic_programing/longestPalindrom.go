package dynamicprograming

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		c1 := palindromHeight(s, i, i)
		c2 := palindromHeight(s, i, i+1)
		var length int
		if c1 > c2 {
			length = c1 
		} else {
			length = c2
		}
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start:end + 1]
}

func palindromHeight(s string, start int, end int) int {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	return end - start - 1
}
