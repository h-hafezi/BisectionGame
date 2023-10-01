package MerkleMountainRange

import "crypto/sha256"

// In the MMR file we use a hash function, we make it generic in a function and initialise with Sha256

func Hash(string *[]byte) [32]byte {
	return sha256.Sum256(*string)
}
