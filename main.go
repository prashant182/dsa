package main

import (
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
	bt := bt.NewBinaryTree(node1)
	node1.Left = node2
	node1.Right = node3
	bt.Print()
	/*
		Let's add Node 4 as the child of Node 2 in this tree Making it
					Node(1)
					/    \
			 Node(2)      Node(3)
			   /
			  /
			Node(4)
	*/
	bt.AddNode(4)
	bt.Print()
	/*
		Let's remove Node 3 From this tree; we will end up with the following tree.
		#I have followed this algorithm for this implementation.
		## https://www.geeksforgeeks.org/deletion-binary-tree/
					Node(1)
					/    \
			 Node(2)      Node(4)
	*/
	bt.DeleteNode(3)
	bt.Print()
}
