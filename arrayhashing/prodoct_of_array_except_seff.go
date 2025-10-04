// medium
package arrayhashing

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	for i := range nums {
		res[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	suffix := 1

	for i := len(nums) - 2; i >= 0; i-- {
		suffix = suffix * nums[i+1]
		res[i] = res[i] * suffix
	}

	return res
}
