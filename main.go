package main

import (
	"Merkle_Mountain_Range/MerkleMountainRange"
)

func main() {
	/*
		// define a message
		// initialise a key
		var key1 Signature.Key
		key1.NewKey()
		var key2 Signature.Key
		key2.NewKey()
		var key3 Signature.Key
		key3.NewKey()
		array1 := [3]*Signature.PublicKey{
			key1.GetPublicKey(), key2.GetPublicKey(), key3.GetPublicKey(),
		}
		var message Signature.Message
		stream := []byte("Here is a string....")
		message.InitialiseFromRawBytes(&stream)
		sig1 := key1.GetPrivateKey().Sign(&message)
		sig2 := key2.GetPrivateKey().Sign(&message)
		sig3 := key3.GetPrivateKey().Sign(&message)
		array2 := [3]*Signature.Signature{
			sig1, sig2, sig3,
		}
		aggSig := Signature.AggregateSignatures(array2[:], array1[:])
		epoch := Epoch{signature: aggSig, validators: array1[:]}

		var key11 Signature.Key
		key11.NewKey()
		var key22 Signature.Key
		key22.NewKey()
		var key33 Signature.Key
		key33.NewKey()
		array11 := [3]*Signature.PublicKey{
			key11.GetPublicKey(), key22.GetPublicKey(), key33.GetPublicKey(),
		}
		// makes it into bytes and make all the previous keys sing it
		byteStream1 := Signature.MultiplePublicKeysToByte(array11[:])
		// transform it into a message
		var message1 Signature.Message
		message1.InitialiseFromRawBytes(byteStream1)
		// sign it by each key separately
		sig11 := key1.GetPrivateKey().Sign(&message1)
		sig22 := key2.GetPrivateKey().Sign(&message1)
		sig33 := key3.GetPrivateKey().Sign(&message1)
		array22 := [3]*Signature.Signature{
			sig11, sig22, sig33,
		}
		aggSig1 := Signature.AggregateSignatures(array22[:], array1[:])
		epoch1 := Epoch{signature: aggSig1, validators: array11[:]}
		fmt.Println(epoch1.IsEpochValid(epoch))
	*/

	var mmr MerkleMountainRange.MerkleMountainRange
	mmr.New()
	e1 := []byte("e1")
	mmr.AppendElement(&e1)
	e2 := []byte("e2")
	mmr.AppendElement(&e2)
	e3 := []byte("e3")
	mmr.AppendElement(&e3)
	e4 := []byte("e4")
	mmr.AppendElement(&e4)
	e5 := []byte("e5")
	mmr.AppendElement(&e5)
	e6 := []byte("e6")
	mmr.AppendElement(&e6)
	e7 := []byte("e7")
	mmr.AppendElement(&e7)
}
