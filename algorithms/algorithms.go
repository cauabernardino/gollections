package algorithms

import (
	"fmt"

	"github.com/cauabernardino/gollections/algorithms/hashes"
)

func Run() {
	run_hashes()
}

func run_hashes() {
	// 01
	nums1 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	ans1 := hashes.ContainsDuplicate(nums1)
	fmt.Println(ans1)

	// 02
	s := "anagram"
	t := "nagaram"
	ans2 := hashes.IsAnagram(s, t)
	fmt.Println(ans2)

	// 03
	nums3 := []int{5, 3, 2, 11, 7, 15}
	target := 9
	ans3 := hashes.TwoSum(nums3, target)
	fmt.Println(ans3)

	// 04
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans4 := hashes.GroupAnagrams(strs)
	fmt.Println(ans4)

	// 05
	nums5 := []int{1, 1, 1, 2, 3, 3, 3, 4, 5, 2, 2, 2, 3}
	k := 2
	ans5 := hashes.TopKFrequent(nums5, k)
	fmt.Println(ans5)

	// 06
	vals6 := []int{1, 2, 3, 4}
	ans6 := hashes.ProductExceptSelf(vals6)
	fmt.Println(ans6)

	// 07
	// vals7 := []int{100, 4, 200, 1, 3, 2}
	vals72 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	ans7 := hashes.LongestConsecutive(vals72)
	fmt.Println(ans7)
}
