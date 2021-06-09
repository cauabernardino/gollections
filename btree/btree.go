package btree

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

type Tree struct {
	Root *Node
}

// New creates a new Binary Tree
func New() *Tree {
	return &Tree{}
}

// Insert handles the insertion in Tree level
func (t *Tree) Insert(data int) *Tree {
	if t.Root == nil {
		t.Root = &Node{Data: data, Left: nil, Right: nil}
	} else {
		t.Root.insert(data)
	}

	return t
}

// insert handles the insertion on node level
func (n *Node) insert(data int) {
	if n == nil {
		return
	}

	if data <= n.Data {
		if n.Left == nil {
			n.Left = &Node{Data: data, Left: nil, Right: nil}
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Data: data, Left: nil, Right: nil}
		} else {
			n.Right.insert(data)
		}
	}
}

// PrintInOrder prints the tree values in crescent order
func PrintInOrder(n *Node) {
	if n == nil {
		return
	}

	PrintInOrder(n.Left)
	fmt.Printf("%d ", n.Data)
	PrintInOrder(n.Right)
}

// PrintLevelOrder prints the tree values in order by tree level
func PrintLevelOrder(root *Node) {
	var (
		stack []*Node
		index int
	)

	stack = append(stack, root)
	index = -1

	for {
		index++

		if index == len(stack) {
			fmt.Printf("\n")
			break
		}

		fmt.Printf("%d ", stack[index].Data)

		if stack[index].Left != nil {
			stack = append(stack, stack[index].Left)
		}

		if stack[index].Right != nil {
			stack = append(stack, stack[index].Right)
		}

	}
}

// // PrintLevelOrderWithDeque prints the tree values in order by tree level
// func PrintLevelOrderWithDeque(root *Node) {
// 	dq := deque.New()

// 	dq.Append(root)

// 	for {
// 		if dq.Len() == 0 {
// 			fmt.Printf("\n")
// 			break
// 		}

// 		currInterface, _ := dq.PopLeft()

// 		curr := currInterface.(*Node)

// 		fmt.Printf("%d ", curr.Data)

// 		if curr.Left != nil {
// 			dq.Append(curr.Left)
// 		}

// 		if curr.Right != nil {
// 			dq.Append(curr.Right)
// 		}
// 	}
// }
