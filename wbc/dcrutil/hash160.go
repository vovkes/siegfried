// Copyright (c) 2013-2014 The btcsuite developers
// The WBC developers. Copyright (c) 2017 
//

package dcrutil

import (
	"hash"

	"golang.org/x/crypto/ripemd160"

	"github.com/wbcoin/wbc/chaincfg/chainhash"
)

// Calculate the hash of hasher over buf.
func calcHash(buf []byte, hasher hash.Hash) []byte {
	hasher.Write(buf)
	return hasher.Sum(nil)
}

// Hash160 calculates the hash ripemd160(hash256(b)).
func Hash160(buf []byte) []byte {
	return calcHash(chainhash.HashB(buf), ripemd160.New())
}
