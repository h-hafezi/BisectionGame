package Signature

import (
	"crypto/rand"
	bls "github.com/ethereum/go-ethereum/crypto/bls12381"
	"math/big"
)

// defining tha global parameters which will be accessed by all functions
var engine *bls.Engine = bls.NewPairingEngine()

type Key struct {
	privateKey *PrivateKey
	publicKey  *PublicKey
}

func (key *Key) NewKey() {
	// Max value, a 130-bits integer, i.e 2^130 - 1
	var max *big.Int = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(130), nil)
	// Generate cryptographically strong pseudo-random between [0, max)
	n, _ := rand.Int(rand.Reader, max)
	c := engine.G2.One()
	engine.G2.MulScalar(c, engine.G2.One(), n)
	// setting the attributes
	key.privateKey = &PrivateKey{n}
	key.publicKey = &PublicKey{c}
}

func (key *Key) GetPublicKey() *PublicKey {
	return key.publicKey
}

func (key *Key) GetPrivateKey() *PrivateKey {
	return key.privateKey
}
