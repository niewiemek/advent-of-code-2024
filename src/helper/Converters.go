package helper

import "strconv"

func ToIntArray(stringArray []string) (nums []int) {
	var _ error
	nums = make([]int, len(stringArray), len(stringArray))
	for i, value := range stringArray {
		nums[i], _ = strconv.Atoi(value)
	}

	return nums
}
