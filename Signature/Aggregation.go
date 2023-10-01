package Signature

import (
	"crypto/sha256"
	"encoding/binary"
	"math/big"
)

func AggregateSignatures(signatures []*Signature, publicKeys []*PublicKey) *Signature {
	// raise an error if the length of two arrays aren't equal!
	// first have the public keys to a byte stream
	byteStream := MultiplePublicKeysToByte(publicKeys)
	// hash the byte stream as the seed
	hash := sha256.Sum256(*byteStream)
	// set the initial value as zero
	res := engine.G1.Zero()
	// iterating through all signatures
	for i := 0; i < len(signatures); i++ {
		// index to byte
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, uint32(i))
		// append the current index to hash and hash it
		currentHash := sha256.Sum256(append(hash[:], bs...))
		z := new(big.Int)
		z.SetBytes(currentHash[:])
		// multiply the current signature by the current hash
		currentSig := engine.G1.New()
		engine.G1.MulScalar(currentSig, signatures[i].value, z)
		// add to the sum
		engine.G1.Add(res, res, currentSig)
	}
	return &Signature{value: res}
}

func AggregatePublicKeys(publicKeys []*PublicKey) *PublicKey {
	// raise an error if the length of two arrays aren't equal!
	// first have the public keys to a byte stream
	byteStream := MultiplePublicKeysToByte(publicKeys)
	// hash the byte stream as the seed
	hash := sha256.Sum256(*byteStream)
	// set the initial value as zero
	res := engine.G2.Zero()
	// iterating through all signatures
	for i := 0; i < len(publicKeys); i++ {
		// index to byte
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, uint32(i))
		// append the current index to hash and hash it
		currentHash := sha256.Sum256(append(hash[:], bs...))
		z := new(big.Int)
		z.SetBytes(currentHash[:])
		// multiply the current signature by the current hash
		currentPublicKey := engine.G2.New()
		engine.G2.MulScalar(currentPublicKey, publicKeys[i].value, z)
		// add to the sum
		engine.G2.Add(res, res, currentPublicKey)
	}
	return &PublicKey{value: res}
}

func VerifyAggregatedSignature(publicKeys []*PublicKey, message *Message, signature *Signature) bool {
	aggregatedPublicKey := AggregatePublicKeys(publicKeys)
	return aggregatedPublicKey.Verify(message, signature)
}
