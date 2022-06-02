package hashes

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	pre := 1
	for i := range nums {
		res[i] = pre
		pre *= nums[i]
	}

	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= post
		post *= nums[i]
	}
	return res
}
