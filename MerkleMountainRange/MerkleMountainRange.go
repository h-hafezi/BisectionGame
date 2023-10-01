package MerkleMountainRange

type MMR interface {
	appendElement()
	proveMembership()
}

type MerkleMountainRange struct {
	// The list of all the elements
	elements []*[]byte
	// It's 32 bytes since output of SHA256 is 32 bytes
	Peaks []*Node
}

// New - Initialise a merkle mountain range with no element and hence zero peaks
func (mmr *MerkleMountainRange) New() {
	var elements []*[]byte
	var peaks []*Node
	mmr.elements = elements
	mmr.Peaks = peaks
}

func (mmr *MerkleMountainRange) GetNumberOfElements() int {
	return len(mmr.elements)
}

// GetProofPath : get the proof path, it returns a proof object as well as an integer which is peakIndex of the path
func (mmr *MerkleMountainRange) GetProofPath(index int) (*MMRProof, int) {
	// get the length and make sure the index is lower than the length
	if len(mmr.elements) < index {
		panic("invalid index")
	}
	// find the peak (tree) that hold the index
	var peakIndex int
	// get the peak first
	for i := 0; i < len(mmr.Peaks); i++ {
		if index <= mmr.Peaks[i].index {
			peakIndex = i
			break
		}
	}
	// define the outputs, note that length of path are equal but the last element of path is set to be zero
	var path []*Node
	var direction []int
	// the current node
	currentNode := mmr.Peaks[peakIndex]
	// go down the tree and the proof path
	for {
		path = append(path, currentNode)
		// check we need to go left or right
		if currentNode.index == index {
			// add it to the path and terminate
			direction = append(direction, 0)
			break
		} else if currentNode.Left.index < index {
			// we need to search the right tree
			currentNode = currentNode.Right
			// add the direction to the array
			direction = append(direction, 1)
		} else {
			// we need to search the left tree
			currentNode = currentNode.Left
			// add the node and direction to the array
			direction = append(direction, -1)
		}
	}
	return &MMRProof{direction: direction, path: path}, peakIndex
}

func (mmr *MerkleMountainRange) AppendElement(element *[]byte) {
	// update the list of elements
	mmr.elements = append(mmr.elements, element)
	// if the length is one add the first peak to the array
	if len(mmr.elements) == 1 {
		temp := Hash(element)
		node := Node{1, element, &temp, nil, nil, nil, 1}
		mmr.Peaks = append(mmr.Peaks, &node)
	} else {
		// an element is either added as a single root or merges two already built trees
		length := len(mmr.Peaks)
		n := len(mmr.elements)
		if isMerge(n) {
			// concat = element + peak[n-1] + peak[n-2]
			concat := append((*mmr.Peaks[length-1]).hash[:], (*mmr.Peaks[length-2]).hash[:]...)
			concat = append(*element, concat...)
			// temp = Hash(concat)
			temp := Hash(&concat)
			// introduce the new node
			node := Node{n, element, &temp, mmr.Peaks[length-2], mmr.Peaks[length-1], nil, 2*mmr.Peaks[length-2].childrenNumber + 1}
			// updating parents
			mmr.Peaks[length-2].parent = &node
			mmr.Peaks[length-1].parent = &node
			// popping the last two elements
			mmr.Peaks = mmr.Peaks[:length-2]
			// adding the new node to the list of peaks
			mmr.Peaks = append(mmr.Peaks, &node)
		} else {
			n := len(mmr.elements)
			temp := Hash(element)
			node := Node{n, element, &temp, nil, nil, nil, 1}
			mmr.Peaks = append(mmr.Peaks, &node)
		}
	}
}

// return the largest integer in form of 2^k-1 less than n
func largestPowerOfTwo(n int) int {
	res := 1
	for {
		if res == n || res-1 == n {
			return res - 1
		} else if res > n {
			return res/2 - 1
		} else {
			res *= 2
		}
	}
}

// GetPeakSizes : get the correct size of peaks for size n
func GetPeakSizes(n int) []int {
	if n == 0 {
		result := [0]int{}
		return result[:]
	}
	if n == 1 {
		result := [1]int{1}
		return result[:]
	} else {
		k := largestPowerOfTwo(n)
		res1 := [1]int{k}
		res2 := GetPeakSizes(n - k)
		res := append(res1[:], res2...)
		return res
	}
}

// if the element is added merges two least peaks or add a new peak
func isMerge(n int) bool {
	for {
		if n == 1 || n == 2 {
			return false
		}
		if n == 3 {
			return true
		}
		m := largestPowerOfTwo(n)
		if m == n {
			return true
		} else {
			n -= m
		}
	}
}
