package plusOne

func plusOne(digit []int) []int {
	var result = make([]int, len(digit)+1, len(digit)+1)
	offset := 1
	for index := len(digit) - 1; index >= 0; index-- {
		if digit[index]+offset > 9 {
			result[index+1] = 0
			offset = 1
		} else {
			result[index+1] = digit[index] + offset
			offset = 0
		}
	}
	if offset == 1 {
		result[0] = 1
	} else {
		result = result[1:]
	}
	return result
}
