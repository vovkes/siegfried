// Copyright (c) 2013-2014 The btcsuite developers
// Copyright (c) 2015 The WBC developers
//

/*
This test file is part of the btcutil package rather than than the
btcutil_test package so it can bridge access to the internals to properly test
cases which are either not possible or can't reliably be tested via the public
interface. The functions are only exported while the tests are being run.
*/

package dcrutil

import (
	"bitbucket.org/siegfriedvmblockchain/siegfried/base58"
	"bitbucket.org/siegfriedvmblockchain/siegfried/wbc/chaincfg/chainec"

	"golang.org/x/crypto/ripemd160"
)

// SetBlockBytes sets the internal serialized block byte buffer to the passed
// buffer.  It is used to inject errors and is only available to the test
// package.
func (b *Block) SetBlockBytes(buf []byte) {
	b.serializedBlock = buf
}

// TstAppDataDir makes the internal appDataDir function available to the test
// package.
func TstAppDataDir(goos, appName string, roaming bool) string {
	return appDataDir(goos, appName, roaming)
}

// TstAddressPubKeyHash makes an AddressPubKeyHash, setting the
// unexported fields with the parameters hash and netID.
func TstAddressPubKeyHash(hash [ripemd160.Size]byte,
	netID [2]byte) *AddressPubKeyHash {
	return &AddressPubKeyHash{
		hash:  hash,
		netID: netID,
	}
}

// TstAddressScriptHash makes an AddressScriptHash, setting the
// unexported fields with the parameters hash and netID.
func TstAddressScriptHash(hash [ripemd160.Size]byte,
	netID [2]byte) *AddressScriptHash {

	return &AddressScriptHash{
		hash:  hash,
		netID: netID,
	}
}

// TstAddressPubKey makes an AddressPubKey, setting the unexported fields with
// the parameters.
func TstAddressPubKey(serializedPubKey []byte, pubKeyFormat PubKeyFormat,
	netID [2]byte) *AddressSecpPubKey {

	pubKey, _ := chainec.Secp256k1.ParsePubKey(serializedPubKey)
	return &AddressSecpPubKey{
		pubKeyFormat: pubKeyFormat,
		pubKey:       chainec.PublicKey(pubKey),
		pubKeyHashID: netID,
	}
}

// TstAddressSAddr returns the expected script address bytes for
// P2PKH and P2SH wbcoin addresses.
func TstAddressSAddr(addr string) []byte {
	decoded := base58.Decode(addr)
	return decoded[2 : 2+ripemd160.Size]
}
