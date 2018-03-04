// Copyright (c) 2013-2014 The btcsuite developers
// The WBC developers. Copyright (c) 2017 
//

package dcrutil

const (
	// AtomsPerCent is the number of atomic units in one coin cent.
	AtomsPerCent = 1e6

	// AtomsPerCoin is the number of atomic units in one coin.
	AtomsPerCoin = 1e8

	// MaxAmount is the maximum transaction amount allowed in atoms.
	// WBC - Changeme for release
	MaxAmount = 21e6 * AtomsPerCoin
)
