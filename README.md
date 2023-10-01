# BisectionGame

In order to implement the bisection game from (["Proofs of Proof-of-Stake with Sublinear Complexity"](https://arxiv.org/pdf/2209.08673.pdf)) for a permissioned blockchain (e.g. Quorum), we reimplement the code in golang, our code, in brief, consists of the following components.

## Signature

here we implement a simple BLS library based on BLS123-81, supporting the signature/public key aggregation, serialisation/deserialisation of public key and signatures etc.

## Merkle Mountain Range

Merkle mountain range is a similar concept to Merkle root but is more efficient for append-only purposes. You may find out more about it (["here"](https://docs.grin.mw/wiki/chain-state/merkle-mountain-range/)). We implement functionalities such as AppendELement, GetProofPath and VerifyProof. In more details, this folder consists of the following files.


> Node

This struct is used to build the trees.
```
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
```
> MMR

```
type MerkleMountainRange struct {
	// The list of all the elements
	elements []*[]byte
	// It's 32 bytes since output of SHA256 is 32 bytes
	Peaks []*Node
}
```
## Epoch

Epoch is a generic data structure in the following format:

```
type Epoch struct {
	// signers, the order is important
	validators []*Signature.PublicKey
	// Aggregated Signature of the previous validators on the current validators
	signature *Signature.Signature
}
```
We implement the following functionalities for an Epoch:
```
func (epoch *Epoch) IsEpochValid(previous Epoch) bool {
    // checks if one epoch is valid given it's the previous epoch
 }

```
We also implement Encode/Decode on an Epoch in order for serialisation/deserialisation.

## Bisection

The bisection game is defined on two MMRs (note it can be implemented on two-byte streams representing MMRs too). The bisection game is implemented generically as follows:

```
func bisectionGame(mmr [2]*MerkleMountainRange.MerkleMountainRange) int {
  // checking size, first check is the two MMRs are of the same size
  	
  // get the index of the first different peak and making sure that the two chains are not equal
  	
  // set the currents nodes to root of currrent peaks and update these values as traversing the tree
  
  // going down both the trees
  for {
      // the basis case where the node has no children
      if currentNode1.GetChildren() == 1 {
      // get the previous nodes of the current ones and check with the Eval function which one is correct
      }
      // there are three different cases one
      1- the left children are inconsistent ==> update current nodes to the left children
      2- the right children are inconsistent ==> update current nodes to the right children
      3- both right and left children are consistent ==> current nodes have inconsistent value ==> run Eval on them 
  }
```
This functionality implicitly takes a boolean function as an argument Eval which evaluates two consecutive Epochs. The output of bisection game is as follows:

$$\begin{cases}1&:\mbox{the first MMR is incorrect}\\
2&:\mbox{the second MMR is incorrect}\\
3&:\mbox{both MMRs are incorrect}
\end{cases}$$

Note that if one MMR is incorrect, it doesn't necessaily mean the other one is correct
