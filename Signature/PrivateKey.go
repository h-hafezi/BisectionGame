package Signature

import "math/big"

type PrivateKey struct {
	value *big.Int
}

func (privateKey *PrivateKey) Sign(message *Message) *Signature {
	c := engine.G1.One()
	sig := engine.G1.MulScalar(c, message.hash, privateKey.value)
	return &Signature{value: sig}
}
