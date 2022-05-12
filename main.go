package main

import "fmt"

/*
BOOGABLOCK
	DONE:
		serialization, validation, mining, blockchain base implementation
	TODO:
	Transactions (hashmerkleroot)
		hashmerkleroot is updated when a transaction is made (store transacations)
	Network
		network must be created to allow transactions to begin with, p2p transactions from a central db
		addresses must be handled here, randomly generated addresses assigned to new users in the chain
	Database
		redis as the database of choice, smaller scale project so redis's speed and ease of use is nice
		need to serialize the blocks, use protocol buffers to serialize and decode
*/

func main() {
	myChain := NewChain("GENESIS BLOCK")

	// grabbing the genesis block from a new chain
	gen_val := myChain.Database.rdb.Get(myChain.Database.ctx, "1")
	if gen_val.Err() != nil {
		fmt.Println(gen_val.Err())
	}
	fmt.Println(gen_val.Val())

	myChain.AddBlock("HEYO THIS IS THE FIRST ADDED BLOCK")

	// grabbing the new block added to a chain
	added_val := myChain.Database.rdb.Get(myChain.Database.ctx, "1")
	if added_val.Err() != nil {
		fmt.Println(added_val.Err())
	}
	fmt.Println(added_val.Val())

}
