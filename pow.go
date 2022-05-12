package main

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

// proof of work struct (has target and block desired)
type ProofOfWork struct {
	powBlock *Block
	target   *big.Int
}

// generate new pow
func NewPoW(b *Block) *ProofOfWork {
	// target is upper boundary of range, if # is lower than boundary then its valid
	target := big.NewInt(1)
	// bitshifting
	target.Lsh(target, HASH_SIZE-TARGET_BITS)
	return &ProofOfWork{b, target}
}

// hash data
func (PoW *ProofOfWork) prepData(nonce int) []byte {
	// FLATTEN BYTE ARRAYS - IMPORTANT
	return bytes.Join(
		[][]byte{
			PoW.powBlock.PrevHash,
			PoW.powBlock.Data,
			Int64ToByte(PoW.powBlock.TimeStamp),
			Int64ToByte(int64(TARGET_BITS)),
			Int64ToByte(int64(nonce)),
		},
		[]byte{},
	)
}

// mine
func (PoW *ProofOfWork) Mine() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for {
		data := PoW.prepData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(PoW.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

func (PoW *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := PoW.prepData(PoW.powBlock.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(PoW.target) == -1
}
