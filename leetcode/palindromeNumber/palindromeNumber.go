package palindromeNumber

import (
	"math"
)

func isPalindrome(x int) bool {
	if x > math.MaxInt64 {
		return false
	}
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	var seq []int
	for x > 0 {
		val := x % 10
		seq = append(seq, val)
		x = x / 10
	}
	for idx, val := range seq {
		if val == seq[len(seq)-idx-1] {
			if idx > len(seq)/2+1 {
				return true
			}
			continue
		} else {
			return false
		}
	}
	return true
}
