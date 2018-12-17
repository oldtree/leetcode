package monotonicArray

/*
如果数组是单调递增或单调递减的，那么它是单调的。
如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。 如果对于所有 i <= j，A[i]> = A[j]，那么数组 A 是单调递减的。
当给定的数组 A 是单调数组时返回 true，否则返回 false。
*/

func IsMonotonic(array []int) bool {
	mark := 0
	for index, value := range array {
		for idx := index; idx < len(array); idx++ {
			if index <= idx && value <= array[idx] {
			} else {
				mark++
				break
			}
		}
		if mark != 0 {
			break
		}
	}
	if mark == 0 {
		goto SUCCESS
	}
	mark = 0
	for index, value := range array {
		for idx := index; idx < len(array); idx++ {
			if index <= idx && value >= array[idx] {
			} else {
				mark++
				break
			}
		}
		if mark != 0 {
			break
		}
	}
	if mark == 0 {
		goto SUCCESS
	} else {
		goto FAILED
	}

SUCCESS:
	return true
FAILED:
	return false

}
