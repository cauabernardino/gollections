package hashes

import "reflect"

// 02 isAnagram - needs package reflect
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s_hash := make(map[rune]int)
	t_hash := make(map[rune]int)

	for _, c := range s {
		s_hash[c]++
	}
	for _, c := range t {
		t_hash[c]++
	}

	return reflect.DeepEqual(s_hash, t_hash)
}
