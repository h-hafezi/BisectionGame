package MerkleMountainRange

import (
	"encoding/hex"
	"fmt"
)

type Node struct {
	// index of the node
	index int
	// the data
	data *[]byte
	// hash of the data in the node
	hash *[32]byte
	// the children
	Left  *Node
	Right *Node
	// parent node
	parent *Node
	// number of children (including itself) which is always equal to 2^k-1
	childrenNumber int
}

func (node *Node) PrintNode() string {
	return fmt.Sprintf("Index: %d, data: %s, hash: %s, left-node (adr): %p,"+
		"right-node (adr): %p, parent (adr): %p", node.index, string(*node.data), hex.EncodeToString((*node.hash)[:]),
		node.Left, node.Right, node.parent)
}

func (node *Node) GetData() *[]byte {
	return node.data
}

func (node *Node) GetHash() *[32]byte {
	return node.hash
}

func (node *Node) GetSibling() *Node {
	parent := node.parent
	if parent == nil {
		panic("no sibling exists")
	}
	if parent.Left == node {
		return parent.Right
	} else {
		return parent.Left
	}
}

func (node *Node) GetChildren() int {
	return node.childrenNumber
}
