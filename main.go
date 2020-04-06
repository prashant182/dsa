package main

import (
	"fmt"

	bt "github.com/prashant182/dsa/binarytree"
)

func main() {

	/*
		Creating a Tree with Node structure like this.
					Node(1)
					/    \
			 Node(2)      Node(3)
	*/
	node1 := bt.NewNode(1)
	node2 := bt.NewNode(2)
	node3 := bt.NewNode(3)
	bt1 := bt.NewBinaryTree(node1)
	node1.Left = node2
	node1.Right = node3
	bt1.Print()
	/*
		Let's add Node 4 as the child of Node 2 in this tree Making it
					Node(1)
					/    \
			 Node(2)      Node(3)
			   /
			  /
			Node(4)
	*/
	bt1.AddNode(4)
	bt1.Print()
	/*
		Let's remove Node 3 From this tree; we will end up with the following tree.
		#I have followed this algorithm for this implementation.
		## https://www.geeksforgeeks.org/deletion-binary-tree/
					Node(1)
					/    \
			 Node(2)      Node(4)
	*/
	bt1.DeleteNode(3)
	bt1.Print()

	/*
		Let's check if the above tree is continuous or not
					Node(4)
					/    \
			 Node(5)      Node(3)
			 /
		 Node(6)

	*/
	node4 := bt.NewNode(4)
	node5 := bt.NewNode(5)
	node3 = bt.NewNode(3)
	node6 := bt.NewNode(6)
	node4.Left = node5
	node5.Left = node6
	node4.Right = node3
	ct1 := bt.NewBinaryTree(node4)
	fmt.Println("IsContinuous :", ct1.IsContinuous())
	/*
		Let's check how different traversal algorithms works for above tree
		InOrder := 6,5,4,3
		PreOrder := 4,5,6,3
		PostOrder := 6,5,3,4
	*/
	ct1.InOrder()
	ct1.PreOrder()
	ct1.PostOrder()

	/*
		Array Based binary tree example
	*/
	at := bt.NewArrayBT(10)
	at.SetRoot("A")
	at.SetLeft("C", 0)
	at.SetRight("D", 0)
	at.Print()

}
