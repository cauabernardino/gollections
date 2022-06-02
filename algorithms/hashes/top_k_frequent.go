package hashes

func TopKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}

	length := len(nums)

	count := make(map[int]int)
	freq := make([][]int, length+1)

	for _, n := range nums {
		count[n]++
	}

	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	var res []int
	for i := len(freq) - 1; i > 0; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}

	return nil
}
