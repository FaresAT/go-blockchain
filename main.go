package main

import "go-blockchain/cmd"

/*
BOOGABLOCK
	DONE:
		blockchain base implementation, serialization, validation, mining, redis database
	TODO:
	Transactions (hashmerkleroot)
		hashmerkleroot is updated when a transaction is made (store transactions)
	Network
		network must be created to allow transactions to begin with, p2p transactions from a central db
		addresses must be handled here, randomly generated addresses assigned to new users in the chain
*/

func main() {
	cmd.Execute()
}
