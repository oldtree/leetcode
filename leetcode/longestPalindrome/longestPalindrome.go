package longestPalindrome

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
Example 1:
	Input: "babad"
	Output: "bab"
	Note: "aba" is also a valid answer.
Example 2:
	Input: "cbbd"
	Output: "bb"

*/

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	maxLength := 0
	maxStart := 0
	tempMaxLength := 0
	tempMaxStart := 0
	for index, _ := range s { //迭代遍历
		for idx, val := range s[index:] {
			mark := 0
			tempMaxLength = 0
			tempMaxStart = index
			for id := index; id < index+idx+1; id++ {
				if rune(s[id]) == val {
					mark++
					if mark == 2 {
						tempMaxLength++
						goto LOOPEND
					}
					tempMaxStart = id
					tempMaxLength++
				} else {
					if mark == 1 {
						tempMaxLength++
					}
				}
			}
			if mark == 1 {
				tempMaxLength = 0
			}
		}
	LOOPEND:
		if tempMaxLength > maxLength {
			maxLength = tempMaxLength
			maxStart = tempMaxStart
		}
	}
	if maxLength == 0 {
		return string(s[0])
	}
	return s[maxStart : maxStart+maxLength]
}
