package main

import "fmt"

/*
BOOGABLOCK
	DONE:
		serialization, validation, mining, database, blockchain base implementation
	TODO:
	Transactions (hashmerkleroot)
		hashmerkleroot is updated when a transaction is made (store transactions)
	Network
		network must be created to allow transactions to begin with, p2p transactions from a central db
		addresses must be handled here, randomly generated addresses assigned to new users in the chain
	Database
		redis as the database of choice, smaller scale project so redis's speed and ease of use is nice
		need to serialize the blocks, use protocol buffers to serialize and decode
*/

func main() {
	myChain := NewChain("GENESIS BLOCK")
	myChain.AddBlock("HEYO THIS IS THE FIRST ADDED BLOCK")

	var b *Block

	t := myChain.Database.rdb.Get(myChain.Database.ctx, "1")
	//fmt.Println(t.Val())
	b, _ = DeserializeBlock([]byte(t.Val()))
	fmt.Println(string(b.Data))
}
