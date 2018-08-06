package searchalgorithms

// TODO: implement a binary search tree;
// Example:
// 				5
//			  /   \
//			9      11
//         / \    /  \
//        3   5  7    2
//              /
//			  (6)
//
// Methods:
// Add: adds a new node to the tree according to the following rule: new nodes
// SMALLER than the parent go to the left, new nodes LARGER or IDENTICAL to
// the parent go the right; this procedure is done recursively until a parent
// without a child in the particular position is found and the new node will
// be inserted in that position; e.g.: tree.Add(6) --> right --> left --> left

// Tree implements a binary search tree with a certain value at its root and
// right and left children (both stored in the root node)
type Tree struct {
	root  *Node
	depth int
}

// Node is a node in a binary search tree
type Node struct {
	leftChild  *Node
	rightChild *Node
	data       interface{}
}

// Add adds a new node to a binary search tree
func (t *Tree) Add(data interface{}) {
	// TODO: implement
}

// Delete removes the deepest node of a binary search tree that holds a certain data
// value and returns true if a value was removes;  a parent node cannot be removes
// with this method
func (t *Tree) Delete(Data interface{}) bool {
	// TODO: implement
}

// Depth returns the depth of a binary search tree
func (t *Tree) Depth() int {
	return t.depth
}
