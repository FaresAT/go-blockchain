package main

import (
	"encoding/binary"
)

// difficulty (target bits, upper)
const TARGET_BITS = 24

// hash size
const HASH_SIZE = 256

func Int64ToByte(convert_int int64) []byte {
	convert_byte := make([]byte, 8)
	binary.BigEndian.PutUint64(convert_byte, uint64(convert_int))
	return convert_byte
}
