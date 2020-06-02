package addStrings

func addStrings(num1 string, num2 string) string {
	idx1 := len(num1) - 1
	idx2 := len(num2) - 1
	var result []byte
	reIdx := 0
	if idx2 > idx1 {
		result = make([]byte, idx2+1)
		reIdx = idx2
	} else {
		result = make([]byte, idx1+1)
		reIdx = idx1
	}
	mov := 0
	sub := 0
	end := 0
	val1 := 0
	val2 := 0
	for end != 2 {
		if idx2 < 0 && idx1 < 0 {
			end = 2
			break
		}
		if idx1 >= 0 {
			val1 = int(num1[idx1] - '0')
			idx1--
		} else {
			val1 = 0
		}
		if idx2 >= 0 {
			val2 = int(num2[idx2] - '0')
			idx2--
		} else {
			val2 = 0
		}
		temp := val1 + val2 + mov
		mov = temp / 10
		sub = temp % 10
		result[reIdx] = byte('0' + sub)
		reIdx--
	}
	if mov != 0 {
		result = append([]byte{byte('0' + mov)}, result...)
	}
	return string(result)
}
