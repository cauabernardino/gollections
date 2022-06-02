package hashes

func LongestConsecutive(nums []int) int {
	set := make(map[int]bool)
	longest := 0

	for _, num := range nums {
		set[num] = true
	}

	for _, n := range nums {
		if _, ok := set[n-1]; !ok {
			length := 1

			ok := true
			for ok {
				if _, ok := set[n+length]; ok {
					length++
				} else {
					break
				}
			}

			if length > longest {
				longest = length
			}
		}
	}
	return longest
}
