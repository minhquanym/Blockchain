package main

import (
	"bytes"
	"math/big"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
var base = int64(len(b58Alphabet))

// Encode a byte array to Base58
func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(base)
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)

	for b := range input {
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}

	return result
}

// Decode Base58 data
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	cntZeroBytes := 0

	for b := range input {
		if b == 0x00 {
			cntZeroBytes++
		}
	}

	payload := input[cntZeroBytes:]
	for _, b := range payload {
		charIdx := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, big.NewInt(base))
		result.Add(result, big.NewInt(int64(charIdx)))
	}

	decodedData := result.Bytes()
	decodedData = append(bytes.Repeat([]byte{0x00}, cntZeroBytes), decodedData...)

	return decodedData
}
