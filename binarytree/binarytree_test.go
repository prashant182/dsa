package binarytree

import (
	"testing"
)

func TestBinaryTree_IsSymmetric(t *testing.T) {

	/*
		TestCase 1 :=
								1
						      /   \
							 2     2
							/ \   / \
						   3   4 4   3
	*/
	ns1 := NewNode(1)
	ns2l := NewNode(2)
	ns2r := NewNode(2)
	ns3l := NewNode(3)
	ns3r := NewNode(3)
	ns4l := NewNode(4)
	ns4r := NewNode(4)
	ns1.Left = ns2l
	ns1.Right = ns2r
	ns2l.Left = ns3l
	ns2l.Right = ns4r
	ns2r.Left = ns4l
	ns2r.Right = ns3r
	symmetricTree := NewBinaryTree(ns1)

	/*
		TestCase 2 :=
							1
				   		   / \
						  2   2
						   \   \
						   3    3

	*/
	nt1 := NewNode(1)
	nt2l := NewNode(2)
	nt2r := NewNode(2)
	nt3r := NewNode(3)
	nt3l := NewNode(3)
	nt1.Left = nt2l
	nt1.Right = nt2r
	nt2l.Right = nt3r
	nt2r.Right = nt3l
	asymmetricTree := NewBinaryTree(nt1)

	tests := []struct {
		name string
		tree *BinaryTree
		want bool
	}{
		{
			name: "Symmetric Tree Case 1",
			tree: symmetricTree,
			want: true,
		},
		{
			name: "Asymmetric Tree Case 1",
			tree: asymmetricTree,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.IsSymmetric(); got != tt.want {
				t.Errorf("BinaryTree.IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_PreOrder(t *testing.T) {

	/*
		TestCase 1 :=
								1
						      /   \
							 2     3
							/ \   / \
						   4   5 6   7
	*/
	ns1 := NewNode(1)
	ns2 := NewNode(2)
	ns3 := NewNode(3)
	ns4 := NewNode(4)
	ns5 := NewNode(5)
	ns6 := NewNode(6)
	ns7 := NewNode(7)
	ns1.Left = ns2
	ns1.Right = ns3
	ns2.Left = ns4
	ns2.Right = ns5
	ns3.Left = ns6
	ns3.Right = ns7
	t1 := NewBinaryTree(ns1)
	tests := []struct {
		name string
		tree *BinaryTree
	}{
		{
			name: "PreOrder",
			tree: t1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tree.PreOrder()
		})
	}
}
