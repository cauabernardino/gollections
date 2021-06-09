/*

Gollections

The main package is used just to share some application examples.
To use the actual data structures, import the respective packages.

*/
package main

import (
	"fmt"

	"github.com/cauabernardino/gollections/btree"
)

// Some list of numbers to check outputs
var (
	b1 = []int{1, 14, 3, 7, 4, 5, 15, 6, 13, 10, 11, 2, 12, 8, 9}
	// b2 = []int{1, 2, 5, 3, 6, 4}
	// b3 = []int{
	// 	47, 2, 40, 20, 38, 30, 14, 28, 10, 16, 19, 44, 39, 27, 7, 9, 31,
	// 	12, 43, 21, 5, 41, 34, 49, 13, 33, 3, 4, 25, 22, 29, 15, 32, 35,
	// 	6, 24, 23, 26, 1, 11, 42, 36, 37, 17, 18, 8, 45, 48, 50, 46,
	// }
	// b4 = []int{3, 5, 2, 1, 4, 6, 7}
)

func main() {
	/*
		Binary Search Tree Examples
	*/
	tree := btree.New()

	for _, v := range b1 {
		tree.Insert(v)
	}

	// Traversal Printings
	fmt.Printf("In Order Traversal with Deque:\n")
	btree.PrintInOrder(tree.Root)

	fmt.Printf("\n\nLevel Order Traversal:\n")
	btree.PrintLevelOrder(tree.Root)

	// Number of levels (height)
	height := tree.Height()
	fmt.Printf("\nHeight of tree: %d\n", height)

}
