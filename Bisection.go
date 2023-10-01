package main

import (
	"Merkle_Mountain_Range/MerkleMountainRange"
	"bytes"
)

/*
	It takes two MMRs, 2 elements of these at index i and their previous elements (i-1) and finally proofs of membership
	function eval is used to evaluate consecutive elements, if returns:
		* 1 ==> the first MMR is false
		* 2 ==> the second MMR is false
		* 3 ==> both MMRs are false
*/
func bisectionGame(mmr [2]*MerkleMountainRange.MerkleMountainRange) int {
	// checking size, first check is the two MMRs are of the same size
	if mmr[0].GetNumberOfElements() != mmr[1].GetNumberOfElements() {
		panic("MMRs with inconsistent size")
	}
	// get the size of peaks
	array := MerkleMountainRange.GetPeakSizes(mmr[0].GetNumberOfElements())
	// first find the peak where they're different
	var peakIndex = -1
	for i := 0; i <= len(mmr[0].Peaks); i += 1 {
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
			previous1 := currentNode1.GetPreviousNode()
			previous2 := currentNode2.GetPreviousNode()
			if eval(previous1, currentNode1) == true {
				return 2
			} else if eval(previous2, currentNode2) == true {
				return 1
			} else {
				return 0
			}
		} else {
			// the left are equal so either the right nodes are different so the data at the current node
			if bytes.Equal(currentNode1.Left.GetHash()[:], currentNode2.Left.GetHash()[:]) {
				if bytes.Equal(currentNode1.Right.GetHash()[:], currentNode2.Right.GetHash()[:]) {
					// the current data is different
					previous1 := currentNode1.GetPreviousNode()
					previous2 := currentNode2.GetPreviousNode()
					if eval(previous1, currentNode1) == true {
						return 2
					} else if eval(previous2, currentNode2) == true {
						return 1
					} else {
						return 0
					}
				} else {
					// the right values are different and hence update the current nodes
					currentNode1 = currentNode1.Right
					currentNode2 = currentNode2.Right
				}
			} else {
				// the left values are different and hence update the current nodes
				currentNode1 = currentNode1.Left
				currentNode2 = currentNode2.Left
			}
		}
	}
}

func eval(node1 *MerkleMountainRange.Node, node2 *MerkleMountainRange.Node) bool {
	return true
}
