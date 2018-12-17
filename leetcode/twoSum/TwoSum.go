package twoSum

/*

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
	Given nums = [2, 7, 11, 15], target = 9,
	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
*/

func twoSum(srcSlice []int, target int) []int {
	mark := make([]int, 2, 2)
	for index, value := range srcSlice {
		for idx := index + 1; idx < len(srcSlice); idx++ {
			if value+srcSlice[idx] == target {
				mark[0], mark[1] = index, idx
				break
			}
		}
	}
	return mark
}

func twoSum1(srcSlice []int, target int) []int {
	mark := make([]int, 2, 2)
	for index, value := range srcSlice {
		for idx := index + 1; idx < len(srcSlice); idx++ {
			if value+srcSlice[idx] == target {
				mark[0], mark[1] = index, idx
				break
			}
		}
	}
	return mark
}
