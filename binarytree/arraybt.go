package binarytree

import "fmt"

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

//SetRoot sets the root for the array based binary tree
func (tree *ArrayBT) SetRoot(element string) {
	tree.Elemets[0] = element
}

//NewArrayBT returns the new object for of a type ArrayBT
func NewArrayBT(length int) *ArrayBT {
	return &ArrayBT{Elemets: make([]string, length)}
}

//SetLeft is the method which add the left child for the binary tree
func (tree *ArrayBT) SetLeft(key string, position int) bool {
	t := position*2 + 1
	if tree.Elemets[position] == "" {
		return false
	}
	tree.Elemets[t] = key
	return true
}

//SetRight is the method which sets the right child for the binary tree
func (tree *ArrayBT) SetRight(key string, position int) bool {
	t := position*2 + 2
	if tree.Elemets[position] == "" {
		return false
	}
	tree.Elemets[t] = key
	return true
}

//Print method prints the array bases binary tree
func (tree *ArrayBT) Print() {
	fmt.Println("Tree:")
	for _, elem := range tree.Elemets {
		fmt.Print(elem, ",")
	}
	fmt.Println()
}
