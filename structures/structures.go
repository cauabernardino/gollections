/*

Gollections

The main package is used just to share some application examples.
To use the actual data structures, import the respective packages.

*/
package structures

import (
	"fmt"

	"github.com/cauabernardino/gollections/structures/btree"
	"github.com/cauabernardino/gollections/structures/queue"
	"github.com/cauabernardino/gollections/structures/set"
)

func Run() {
	BinaryTreeExamples()
	SetExamples()
	QueueExamples()
}

/* Binary Search Tree Examples */
func BinaryTreeExamples() {

	// Some list of numbers to check outputs
	var (
		b1 = []int{7, 1, 14, 3, 4, 5, 15, 6, 13, 10, 11, 2, 12, 8, 9}
		// b2 = []int{1, 2, 5, 3, 6, 4}
		// b3 = []int{
		// 	47, 2, 40, 20, 38, 30, 14, 28, 10, 16, 19, 44, 39, 27, 7, 9, 31,
		// 	12, 43, 21, 5, 41, 34, 49, 13, 33, 3, 4, 25, 22, 29, 15, 32, 35,
		// 	6, 24, 23, 26, 1, 11, 42, 36, 37, 17, 18, 8, 45, 48, 50, 46,
		// }
		// b4 = []int{3, 5, 2, 1, 4, 6, 7}
	)

	dataset := b1
	tree := btree.New()

	for _, v := range dataset {
		tree.Insert(v)
	}

	fmt.Printf("---- Binary Search Tree Examples ----\n\n")

	fmt.Printf("Values:\n")
	fmt.Println(dataset)

	// Traversal Printings
	fmt.Printf("\nIn Order Traversal:\n")
	btree.PrintInOrder(tree.Root)

	fmt.Printf("\n\nLevel Order Traversal:\n")
	btree.PrintLevelOrder(tree.Root)

	// Number of levels (height)
	height := tree.Height()
	fmt.Printf("\nHeight of tree: %d\n", height)
}

/* Set Examples */
func SetExamples() {

	mySet := set.New()

	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(3)
	mySet.Add(4)
	mySet.Add(5)

	// Repeated values to show that it won't show twice in Set
	mySet.Add(2)
	mySet.Add(4)

	fmt.Printf("\n---- Set Examples ----\n\n")
	fmt.Println(mySet.Elements())

	n1 := 6
	n2 := 10

	fmt.Printf("%d is element of Set: %v\n", n1, mySet.IsElement(n1))
	fmt.Printf("%d is element of Set: %v\n", n2, mySet.IsElement(n2))

	fmt.Printf("\nAdding the numbers in the set...\n\n")
	mySet.Add(n1)
	mySet.Add(n2)

	fmt.Println(mySet.Elements())
	fmt.Printf("%d is element of Set: %v\n", n1, mySet.IsElement(n1))
	fmt.Printf("%d is element of Set: %v\n", n2, mySet.IsElement(n2))
}

/* Queue Examples*/
func QueueExamples() {
	q := queue.New(5)

	fmt.Printf("\n---- Queue Examples ----\n")
	// Checking the new Queue
	fmt.Printf("\nCreated Queue: %v\n", q.Queue)

	// Adding items
	fmt.Printf("\nAdding values:\n")

	q.Enqueue(1)
	q.Enqueue(2)

	fmt.Printf("\nQueue: %v\n", q.Queue)

	// More items
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Printf("\nQueue: %v\n", q.Queue)

	// // If trying to add more than capacity
	fmt.Printf("\nIf reaches capacity:\n")
	q.Enqueue(6)

	// Checking first and last elements
	fmt.Printf("\nChecking front element of the Queue: ")
	fmt.Println(q.CheckFront())
	fmt.Printf("Checking rear element of the Queue: ")
	fmt.Println(q.CheckRear())

	// Dequeue elements
	fmt.Printf("\nDequeueing 2 elements: ")
	fmt.Println(q.Dequeue(), q.Dequeue())

	fmt.Printf("\nChecking front element of the Queue: ")
	fmt.Println(q.CheckFront())
	fmt.Printf("Checking rear element of the Queue: ")
	fmt.Println(q.CheckRear())

	fmt.Printf("\nQueue: %v\n", q.Queue)

	// Adding new values
	fmt.Printf("\nAdding new value...\n")

	q.Enqueue(6)
	q.Enqueue(7)

	fmt.Printf("\nChecking front element of the Queue: ")
	fmt.Println(q.CheckFront())
	fmt.Printf("Checking rear element of the Queue: ")
	fmt.Println(q.CheckRear())

	fmt.Printf("\nQueue: %v\n", q.Queue)
}
