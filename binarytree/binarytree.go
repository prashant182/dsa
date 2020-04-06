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
			}
			q = append(q, tempNode.Left)
		}
		if tempNode.Right != nil {
			if reflect.DeepEqual(tempNode.Right, deepestNode) {
				tempNode.Right = nil
				return true
			}
			q = append(q, tempNode.Right)
		}
	}
	return false
}

//IsContinuous checks whether the tree is continuous or not. The definition for the continuous tree is derived from here. https://www.geeksforgeeks.org/continuous-tree/
func (tree *BinaryTree) IsContinuous() bool {
	root := tree.Root
	return isContinuous(root)
}

func isContinuous(node *Node) bool {
	if node == nil {
		return true
	}

	if node.Left == nil && node.Right == nil {
		return true
	}

	if node.Right == nil {
		// fmt.Println(node.Data, " - ", node.Left.Data, " = ", Abs(node.Data-node.Left.Data))
		return Abs(node.Data-node.Left.Data) == 1 && isContinuous(node.Left)
	}

	if node.Left == nil {
		// fmt.Println(node.Data, " - ", node.Right.Data, " = ", Abs(node.Data-node.Left.Data))
		return Abs(node.Data-node.Right.Data) == 1 && isContinuous(node.Right)
	}

	// fmt.Println(node.Data, " - ", node.Right.Data, " = ", Abs(node.Data-node.Left.Data))
	// fmt.Println(node.Data, " - ", node.Left.Data, " = ", Abs(node.Data-node.Left.Data))
	return Abs(node.Data-node.Right.Data) == 1 &&
		Abs(node.Data-node.Left.Data) == 1 &&
		isContinuous(node.Left) &&
		isContinuous(node.Right)
}

//IsFoldable is the method which returns whether a tree is foldable or not. Definition https://www.geeksforgeeks.org/foldable-binary-trees/
func (tree *BinaryTree) IsFoldable() bool {

	return false
}

//IsSymmetric is the method which returns whether the given tree is a symmetric tree or not. Definition for the symmetic tree is defined from here. https://www.geeksforgeeks.org/symmetric-tree-tree-which-is-mirror-image-of-itself/
func (tree *BinaryTree) IsSymmetric() bool {
	root := tree.Root
	return isSymmetric(root, root)
}

func isSymmetric(lnode *Node, rnode *Node) bool {
	if lnode == nil && rnode == nil {
		return true
	}

	if lnode != nil && rnode != nil && lnode.Data == rnode.Data {
		return isSymmetric(lnode.Left, rnode.Right) && isSymmetric(lnode.Right, rnode.Left)
	}
	return false
}

//InOrder returns the InOrder traversal path for a given binary tree
func (tree *BinaryTree) InOrder() {
	root := tree.Root
	inOrder(root)
	fmt.Println()
}

func inOrder(node *Node) {
	if (node) == nil {
		return
	}
	inOrder(node.Left)
	fmt.Print(node.Data, " ")
	inOrder(node.Right)
}

//PreOrder returns the PnOrder traversal path for a given binary tree
func (tree *BinaryTree) PreOrder() {
	root := tree.Root
	preOrder(root)
	fmt.Println()
}

func preOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Print(node.Data, " ")
	preOrder(node.Left)
	preOrder(node.Right)
}

//PostOrder returns the PostOrder traversal path for a given binary tree
func (tree *BinaryTree) PostOrder() {
	root := tree.Root
	postOrder(root)
	fmt.Println()
}

func postOrder(node *Node) {
	if node == nil {
		return
	}
	postOrder(node.Left)
	postOrder(node.Right)
	fmt.Print(node.Data, " ")

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

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
