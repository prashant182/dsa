package main

//Node is the type which represents a "Node" in a tree
type Node struct {
	Data  int
	Left  *Node
	right *Node
}

//NewNode returns the "Node" object to be used with a tree
func NewNode(data int) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		right: nil,
	}
}
