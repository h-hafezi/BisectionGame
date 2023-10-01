package main

import (
	"Merkle_Mountain_Range/Signature"
)

/*
Epoch
we build a struct where it has the signing keys and an aggregated signature from the last ocmmittee of singers
to sign the next epoch, all the publicKeys are concatenated and made into a byte array and then the byte array is signed
*/
type Epoch struct {
	// signers, the order is important
	validators []*Signature.PublicKey
	// signature
	signature *Signature.Signature
}

// IsEpochValid
/* Checking an epoch is valid, works as follows:
epoch = {(pk_1, ..., pk_n), sigma}
sigma is an authentic signature on (pk_1, ..., pk_n)
*/
func (epoch *Epoch) IsEpochValid(previous Epoch) bool {
	// transform the list into bytes
	byteStream := Signature.MultiplePublicKeysToByte(epoch.validators)
	// transform it into a message
	var message Signature.Message
	message.InitialiseFromRawBytes(byteStream)
	// get the signature
	return Signature.VerifyAggregatedSignature(previous.validators, &message, epoch.signature)
}

// Encode - encodes the epoch into bytes
func (epoch *Epoch) Encode() []byte {
	var result []byte
	for i := 0; i < len(epoch.validators); i++ {
		// get the key
		publicKey := epoch.validators[i]
		// convert it into byte
		temp := publicKey.PublicKeyToByte()
		result = append(result, *temp...)
	}
	temp := epoch.signature.SignatureToByte()
	result = append(result, *temp...)
	return result
}

// Decode - decodes the epoch from bytes
func (epoch *Epoch) Decode(bytes *[]byte) {
	// initialise the array
	var list []*Signature.PublicKey
	var signature Signature.Signature
	// set the signature
	signature.SignatureFromBytes((*bytes)[len(*bytes)-128:])
	// get the public keys
	remainder := (*bytes)[:len(*bytes)-128]
	for i := 0; i < len(remainder)/256; i++ {
		// get the slice
		temp := (remainder)[i*256 : (i+1)*256]
		// convert it into publicKey
		var publicKey Signature.PublicKey
		publicKey.PublicKeyFromBytes(&temp)
		list = append(list, &publicKey)
	}
	epoch.signature = &signature
	epoch.validators = list
}
