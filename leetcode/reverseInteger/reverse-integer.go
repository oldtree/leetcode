package reverseInteger

/*
Given a 32-bit signed integer, reverse digits of an integer.
Example 1:
	Input: 123
	Output: 321
Example 2:
	Input: -123
	Output: -321
Example 3:
	Input: 120
	Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range:
[−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func reverse(x int) int {
	temp := x
	offset := 0
	ret := 0
	for temp != 0 {
		offset = temp % 10
		ret = ret*10 + offset
		temp = temp / 10
	}
	if ret > (1<<31 - 1) {
		return 0
	} else if ret < (-1 << 31) {
		return 0
	}
	return ret
}
