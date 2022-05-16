package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type BlockChain struct {
	Head     []byte
	Database *redisDB
}

type BCIterator struct {
	Current []byte
	db      *redisDB
}

func (bc *BlockChain) NewIterator() *BCIterator {
	return &BCIterator{bc.Head, bc.Database}
}

func (i *BCIterator) Iterate() *Block {
	var b *Block

	enc := i.db.rdb.Get(i.db.ctx, string(i.Current))
	if enc.Err() == redis.Nil {
		fmt.Println("end of chain")
	} else if enc.Err() != nil {
		panic(enc.Err())
	}
	b, _ = DeserializeBlock([]byte(enc.Val()))
	i.Current = b.PrevHash

	return b
}

func (bc *BlockChain) AddBlock(data string) {
	prev := bc.Database.rdb.Get(bc.Database.ctx, "1")
	new_block := NewBlock(data, []byte(prev.Val()))

	new_sereialized_block, err := new_block.SerializeBlock()
	if err != nil {
		panic(err)
	}

	bc.Database.rdb.Set(bc.Database.ctx, string(new_block.Hash), new_sereialized_block, 0)
	// set new head
	bc.Database.rdb.Set(bc.Database.ctx, "1", new_sereialized_block, 0)
	bc.Head = new_block.Hash
}

func NewChain(genesis string) *BlockChain {

	r := NewRedis()
	genesis_block := GenesisBlock(genesis)
	serialized_genesis, err := genesis_block.SerializeBlock()
	if err != nil {
		panic(err)
	}
	r.rdb.Set(r.ctx, string(genesis_block.Hash), serialized_genesis, 0)
	r.rdb.Set(r.ctx, "1", genesis_block.Hash, 0)

	return &BlockChain{genesis_block.Hash, r}
}
