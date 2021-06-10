/*
Set

- A data structure that store unique values, without any particular order
*/
package set

// Set structure
type Set struct {
	setMap map[int]bool
}

// New initializes a Set
func New() *Set {
	return &Set{}
}

// Add handles the addition of a number in the Set
func (s *Set) Add(n int) {
	if s.setMap == nil {
		s.setMap = make(map[int]bool)
	}

	s.setMap[n] = true
}

// Elements returns the current elements in the Set
func (s *Set) Elements() []int {
	var elements []int

	for k, _ := range s.setMap {
		elements = append(elements, k)
	}

	return elements
}

// IsElement checks if the number is in the Set
func (s *Set) IsElement(n int) bool {
	if s.setMap == nil {
		return false
	}

	return s.setMap[n]

}
