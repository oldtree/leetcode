package regularExpressionMatching

import (
	"strings"
)

/*

Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.

*/
func isMatch(s string, p string) bool {
	srcS := strings.ToLower(s)
	srcP := strings.ToLower(p)
	if srcS == srcP {
		return true
	}
	if srcS == "" && srcP == "" {
		return true
	}
	var (
		stateStart  = 1
		stateDot    = 2
		stateStar   = 3
		stateNormal = 4
		stateEnd    = 5
	)
	var index int

	var state = stateStart
	print(state)
	for _, letter := range srcP {
	CHECK:
		switch {
		case letter >= 97 && letter <= 122:
			state = stateNormal
			if srcS[index] == byte(letter) {
				index++
			}
		case string(letter) == ".":
			state = stateDot
			if srcS[index] <= 'z' && srcP[index] >= 's' {
				continue
			}
		case string(letter) == "*":
			state = stateStar
			if index == 0 {
				index++
			} else {
				if srcS[index] == srcS[index-1] || srcS[index] == 0 {
					index++
					goto CHECK
				}

			}
		default:
			return false
		}
	}
	state = stateEnd
	return true
}
