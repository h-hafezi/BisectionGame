package MerkleMountainRange

import "bytes"

type MMRProof struct {
	// we expect and path and direct have the same number of elements
	path      []*Node
	direction []int
}

func (proof *MMRProof) New(path []*Node, direction []int) {
	if len(path) != len(direction) {
		panic("direction and path of different lengths")
	}
	proof.path = path
	proof.direction = direction
}

func (proof *MMRProof) GetPath() []*Node {
	return proof.path
}

func (proof *MMRProof) GetDirection() []int {
	return proof.direction
}

func (proof *MMRProof) VerifyProof() bool {
	// the basis case
	if len(proof.path) == 1 {
		return true
	}
	// just iterate the path and make sure that the hash relation holds
	for i := 0; i < len(proof.path)-1; i += 1 {
		current := proof.path[i]
		if !VerifySingleStep(current, current.Left, current.Right) {
			return false
		}
	}
	return true
}

// VerifySingleStep : this function takes three simple nodes and makes sure that the hash of father is correct with respect to the children and its data
func VerifySingleStep(parent *Node, left *Node, right *Node) bool {
	concat := append((*right).hash[:], (*left).hash[:]...)
	concat = append(*(*parent).data, concat...)
	res := Hash(&concat)
	return bytes.Equal(res[:], parent.hash[:])
}
