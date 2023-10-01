package Signature

import (
	"crypto/sha256"
	bls "github.com/ethereum/go-ethereum/crypto/bls12381"
	"math/big"
)

type Message struct {
	value *[]byte
	hash  *bls.PointG1
}

func (message *Message) InitialiseFromRawBytes(raw *[]byte) {
	// generate hash from the message to byte
	hash := sha256.Sum256(*raw)
	// byte to big integer
	z := new(big.Int)
	z.SetBytes(hash[:])
	// big integer to pointG1
	c := engine.G1.One()
	engine.G1.MulScalar(c, engine.G1.One(), z)
	// return the new message
	message.value = raw
	message.hash = c
}
