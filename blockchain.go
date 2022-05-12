package main

type BlockChain struct {
	Head     []byte
	Database *redisDB
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

	redis := NewRedis()
	genesis_block := GenesisBlock("GENNY BLOCK")
	serialized_genesis, err := genesis_block.SerializeBlock()
	if err != nil {
		panic(err)
	}
	redis.rdb.Set(redis.ctx, string(genesis_block.Hash), serialized_genesis, 0)
	redis.rdb.Set(redis.ctx, "1", genesis_block.Hash, 0)

	return &BlockChain{genesis_block.Hash, redis}
}
