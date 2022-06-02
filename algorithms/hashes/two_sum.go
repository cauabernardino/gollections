package hashes

// 03 - twoSum
func TwoSum(nums []int, target int) []int {
	sums := make(map[int]int)

	var diff int
	for i, num := range nums {
		diff = target - num
		if _, ok := sums[diff]; ok {
			return []int{sums[diff], i}
		}

		sums[num] = i
	}
	return nil
}
