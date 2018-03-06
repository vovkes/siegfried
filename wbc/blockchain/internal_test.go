// Copyright (c) 2013-2016 The btcsuite developers
// The WBC developers. Copyright (c) 2017 
//

/*
This test file is part of the blockchain package rather than than the
blockchain_test package so it can bridge access to the internals to properly
test cases which are either not possible or can't reliably be tested via the
public interface.  The functions are only exported while the tests are being
run.
*/

package blockchain

import (
	"sort"
	"time"

	"bitbucket.org/siegfriedvmblockchain/siegfried/wbc/chaincfg/chainhash"
	"bitbucket.org/siegfriedvmblockchain/siegfried/wbc/wire"
)

// TstTimeSorter makes the internal timeSorter type available to the test
// package.
func TstTimeSorter(times []time.Time) sort.Interface {
	return timeSorter(times)
}

// TstSetMaxMedianTimeEntries makes the ability to set the maximum number of
// median time entries available to the test package.
func TstSetMaxMedianTimeEntries(val int) {
	maxMedianTimeEntries = val
}

// TstCheckBlockHeaderContext makes the internal checkBlockHeaderContext
// function available to the test package.
func (b *BlockChain) TstCheckBlockHeaderContext(header *wire.BlockHeader, prevNode *blockNode, flags BehaviorFlags) error {
	return b.checkBlockHeaderContext(header, prevNode, flags)
}

// TstNewBlockNode makes the internal newBlockNode function available to the
// test package.
func TstNewBlockNode(blockHeader *wire.BlockHeader, ticketsSpent []chainhash.Hash, ticketsRevoked []chainhash.Hash, voteBits []VoteVersionTuple) *blockNode {
	return newBlockNode(blockHeader, ticketsSpent, ticketsRevoked, voteBits)
}
