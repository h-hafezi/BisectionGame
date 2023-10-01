package main

import (
	"Merkle_Mountain_Range/MerkleMountainRange"
	"bytes"
)

/*
	It takes two MMRs, 2 elements of these at index i and their previous elements (i-1) and finally proofs of membership
	function eval is used to evaluate consecutive elements, it returns the false MMR, note that other isn't necessary correct
*/
func bisectionGame(mmr [2]*MerkleMountainRange.MerkleMountainRange) int {
	// checking size, first check is the two mmr's are of the same size
	if mmr[0].GetNumberOfElements() != mmr[1].GetNumberOfElements() {
		panic("MMRs with inconsistent size")
	}
	// get the size of peaks
	array := MerkleMountainRange.GetPeakSizes(mmr[0].GetNumberOfElements())
	// first find the peak where they're different
	var peakIndex = -1
	for i := len(mmr[0].Peaks) - 1; i >= 0; i -= 1 {
		if !bytes.Equal(mmr[0].Peaks[i].GetHash()[:], mmr[1].Peaks[i].GetHash()[:]) {
			peakIndex = i
		}
		// make sure the sizes are fine
		if mmr[0].Peaks[i].GetChildren() != array[i] {
			return 0
		}
		if mmr[1].Peaks[i].GetChildren() != array[i] {
			return 1
		}
	}
	// making sure the chains are not equal
	if peakIndex == -1 {
		panic("both chains are equal")
	}
	// set the currents nodes
	currentNode1 := mmr[0].Peaks[peakIndex]
	currentNode2 := mmr[1].Peaks[peakIndex]
	// going down both the trees
	for {
		// the basis case where the node has no children
		if currentNode1.GetChildren() == 1 {
			sibling1 := currentNode1.GetSibling()
			sibling2 := currentNode2.GetSibling()
			eval(sibling1, currentNode1)
		} else {
			left1 := currentNode1.Left
			right1 := currentNode1.Right
			left2 := currentNode2.Left
			right2 := currentNode2.Right
			if bytes.Equal(left1.GetHash()[:], left2.GetHash()[:]) {
				if bytes.Equal(right1.GetHash()[:], right2.GetHash()[:]) {
					panic("error")
				} else {
					currentNode1 = right1
					currentNode2 = right2
				}
			} else {
				currentNode1 = left1
				currentNode2 = left2
			}
		}
	}
}

func eval(node1 *MerkleMountainRange.Node, node2 *MerkleMountainRange.Node) bool {
	return true
}
