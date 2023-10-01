package Signature

import bls "github.com/ethereum/go-ethereum/crypto/bls12381"

type PublicKey struct {
	value *bls.PointG2
}

func (publicKey *PublicKey) PublicKeyToByte() *[]byte {
	array := engine.G2.EncodePoint(publicKey.value)
	return &array
}

func (publicKey *PublicKey) PublicKeyFromBytes(array *[]byte) {
	publicKey.value, _ = engine.G2.DecodePoint(*array)
}

func (publicKey *PublicKey) Verify(message *Message, sig *Signature) bool {
	lhs := engine.AddPair(sig.value, engine.G2.One()).Result()
	engine.Reset()
	rhs := engine.AddPair(message.hash, publicKey.value).Result()
	engine.Reset()
	return rhs.Equal(lhs)
}

func MultiplePublicKeysToByte(publicKeyList []*PublicKey) *[]byte {
	var result []byte
	for i := 0; i < len(publicKeyList); i++ {
		// get the key
		publicKey := publicKeyList[i]
		// convert it into byte
		temp := publicKey.PublicKeyToByte()
		result = append(result, *temp...)
	}
	return &result
}

func MultiplePublicKeysFromByte(bytes *[]byte) []*PublicKey {
	// check if the length is authentic
	if len(*bytes)%256 != 0 {
		panic("invalid length")
	}
	// initialise the array
	var list []*PublicKey
	for i := 0; i < len(*bytes)/256; i++ {
		// get the slice
		temp := (*bytes)[i*256 : (i+1)*256]
		// convert it into publicKey
		var publicKey PublicKey
		publicKey.PublicKeyFromBytes(&temp)
		list = append(list, &publicKey)
	}
	return list
}
