package binarytree

//Node is the type which represents a "Node" in a tree
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

//NewNode returns the "Node" object to be used with a tree
func NewNode(data int) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}
