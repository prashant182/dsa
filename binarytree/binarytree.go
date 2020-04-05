package binarytree

import (
	"fmt"
	"reflect"
	"strconv"
)

//BinaryTree is the type of tree
type BinaryTree struct {
	Root *Node
}

//NewBinaryTree is the method which returns the new Node for a binary Tree.
func NewBinaryTree(node *Node) *BinaryTree {
	return &BinaryTree{
		Root: node,
	}
}

//AddNode is the method which add the Node in a given binary Tree.
func (tree *BinaryTree) AddNode(nodeVal int) bool {
	root := tree.Root
	if root == nil {
		newNode := NewNode(nodeVal)
		tree.Root = newNode
		return true
	}
	var q []*Node
	q = append(q, root)
	for len(q) != 0 {
		node := q[0]
		q = q[1:]

		if node.Left == nil {
			node.Left = NewNode(nodeVal)
			return true
		} else {
			q = append(q, node.Left)
		}
		if node.Right == nil {
			node.Right = NewNode(nodeVal)
			return true
		} else {
			q = append(q, node.Right)
		}
	}
	return false
}

//DeleteNode is the method which deletes the node from the tree from https://www.geeksforgeeks.org/deletion-binary-tree/
func (tree *BinaryTree) DeleteNode(nodeVal int) bool {
	var deepestNode *Node
	var nodeToBeDeleted *Node
	root := tree.Root
	var q []*Node
	if root == nil {
		return false
	}

	q = append(q, root)
	for len(q) != 0 {
		tempNode := q[0]
		q = q[1:]
		if tempNode.Data == nodeVal {
			nodeToBeDeleted = tempNode
		}
		deepestNode = tempNode
		if tempNode.Left != nil {
			q = append(q, tempNode.Left)
		}
		if tempNode.Right != nil {
			q = append(q, tempNode.Right)
		}
	}
	nodeToBeDeleted.Data = deepestNode.Data
	/*
		At this point we have addresses for the node to be deleted and the deepestNode which we want to replace it with
		We search for the Node to delete and replace it with the deepestNode
	*/
	q = append(q, root)
	for len(q) != 0 {
		tempNode := q[0]
		q = q[1:]
		if reflect.DeepEqual(tempNode, deepestNode) {
			tempNode = nil
			return true
		}
		if tempNode.Left != nil {
			if reflect.DeepEqual(tempNode.Left, deepestNode) {
				tempNode.Left = nil
				return true
			} else {
				q = append(q, tempNode.Left)
			}
		}
		if tempNode.Right != nil {
			if reflect.DeepEqual(tempNode.Right, deepestNode) {
				tempNode.Right = nil
				return true
			} else {
				q = append(q, tempNode.Right)
			}
		}
	}
	return false
}

//Print method prints the binary tree in a level order traversal
func (tree *BinaryTree) Print() {
	root := tree.Root
	var q []*Node
	if root == nil {
		fmt.Println("NULL Tree")
	}

	fmt.Println("Tree:")
	q = append(q, root)
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		// fmt.Println("Node: ", node)
		fmt.Print(strconv.Itoa(node.Data) + " -> ")
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	fmt.Println()
}

/*
	Binary trees can also be implemented via Arrays. Take a look at this tree for example
				   A(0)
				/		\
			B(1)		C(2)
		  /
	  D(3)
	  We can have
		  Father = p
		  left_child = 2*p + 1
		  right_child = 2*p + 2
		  Tree :=   [A(0), B(1), C(2), D(3)]
		  Indices:= [0,     1,     2,    3 ]
		  father B is on index 1. Left child of B which is D will be on index 3 [2(1)+1 = 3]
*/

//ArrayBT is the Array implementation for a Binary Tree
type ArrayBT struct {
	Elemets []string
}

//SetLeft is the method which add the left child for the binary tree
func (tree *ArrayBT) SetLeft(key string, position int) bool {
	t := position*2 + 1
	if tree.Elemets[position] == "" {
		return false
	} else {
		tree.Elemets[t] = key
	}
	return true
}

//SetRight is the method which sets the right child for the binary tree
func (tree *ArrayBT) SetRight(key string, position int) bool {
	t := position*2 + 1
	if tree.Elemets[position] == "" {
		return false
	} else {
		tree.Elemets[t] = key
	}
	return true

}