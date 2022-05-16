package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	TimeStamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
	Nonce     int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}, 0}
	pow := NewPoW(block)
	nonce, hash := pow.Mine()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func GenesisBlock(genesis string) *Block {
	return NewBlock(genesis, []byte{})
}

func (b *Block) SerializeBlock() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	err := encoder.Encode(b)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DeserializeBlock(buf []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(buf))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
