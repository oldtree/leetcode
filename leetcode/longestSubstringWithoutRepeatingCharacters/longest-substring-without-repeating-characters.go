package longestSubstringWithoutRepeatingCharacters

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:
	Input: "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", which the length is 3.
Example 2:
	Input: "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.
Example 3:
	Input: "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
    Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func LengthOfLongestSubstring(s string) int {
	maxLength := 0
	for index, _ := range s {
		templength := 1
		for idx, val := range s[index+1:] {
			for _, v := range s[index:(index + idx + 1)] {
				if v == val {
					goto COMPUTE
				}
			}
			templength++
		}
	COMPUTE:
		if templength > maxLength {
			maxLength = templength
		}
	}
	return maxLength
}
