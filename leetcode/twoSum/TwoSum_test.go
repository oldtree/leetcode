package twoSum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	result := twoSum(nums[:], 6)
	if result == nil {
		t.Errorf("two sum result error,should be [2,7]")
		return
	}
	t.Logf("two sum %#v", result)
	return
}
