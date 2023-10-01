package Signature

import (
	bls "github.com/ethereum/go-ethereum/crypto/bls12381"
)

type Signature struct {
	value *bls.PointG1
}

func (signature *Signature) SignatureToByte() *[]byte {
	array := engine.G1.EncodePoint(signature.value)
	return &array
}

func (signature *Signature) SignatureFromBytes(array []byte) {
	signature.value, _ = engine.G1.DecodePoint(array)
}
