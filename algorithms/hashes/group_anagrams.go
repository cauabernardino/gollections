package hashes

// 04 - groupAnagrams
func GroupAnagrams(strs []string) [][]string {
	groupsMap := make(map[[26]int][]string)

	for _, str := range strs {
		count := [26]int{}

		for _, c := range str {
			count[c-'a']++
		}
		groupsMap[count] = append(groupsMap[count], str)
	}

	groups := make([][]string, 0)
	for _, v := range groupsMap {
		groups = append(groups, v)
	}

	return groups
}
