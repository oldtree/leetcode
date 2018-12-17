package findMedianSortedArrays

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:
	nums1 = [1, 3]
	nums2 = [2]

	The median is 2.0
Example 2:
	nums1 = [1, 2]
	nums2 = [3, 4]

	The median is (2 + 3)/2 = 2.5
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totlalength := len(nums1) + len(nums2)
	var totalArray []int
	sub1Str := 0
	sub2Str := 0
	for index := 0; index < totlalength; index++ {
		if sub1Str == len(nums1) {
			totalArray = append(totalArray, nums2[sub2Str:]...)
			break
		}
		if sub2Str == len(nums2) {
			totalArray = append(totalArray, nums1[sub1Str:]...)
			break
		}
		if (sub1Str < len(nums1)) && (sub2Str < len(nums2)) && (nums1[sub1Str] >= nums2[sub2Str]) {
			totalArray = append(totalArray, nums2[sub2Str])
			sub2Str++
		} else if (sub1Str < len(nums1)) && (sub2Str < len(nums2)) && (nums1[sub1Str] < nums2[sub2Str]) {
			totalArray = append(totalArray, nums1[sub1Str])
			sub1Str++
		}
	}
	totalArray = totalArray[0:totlalength]
	if len(totalArray)%2 == 0 {
		return float64((totalArray[len(totalArray)/2-1] + totalArray[len(totalArray)/2])) / 2.0
	} else {
		return float64(totalArray[len(totalArray)/2])
	}
}
