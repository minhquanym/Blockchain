package main

import "bytes"

type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// Just extract pubKeyHash from address
// address = 1-byte version + pubKeyHash + 4-byte checksum
// So we only take in range [1:-4]
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// Check whether this output can be used by this pubKey
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}

// Create new transaction output with address and value
func NewTXOutput(value int, address string) *TXOutput {
	result := &TXOutput{value, nil}
	result.Lock([]byte(address))

	return result
}
