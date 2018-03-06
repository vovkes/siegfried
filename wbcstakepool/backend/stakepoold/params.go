// Copyright (c) 2013-2014 The btcsuite developers
// The WBC developers. Copyright (c) 2017 
//

package main

import (
	"bitbucket.org/siegfriedvmblockchain/siegfried/wbc/chaincfg"
	"bitbucket.org/siegfriedvmblockchain/siegfried/wbc/wire"
)

// activeNetParams is a pointer to the parameters specific to the
// currently active wbcoin network.
var activeNetParams = &mainNetParams

// params is used to group parameters for various networks such as the main
// network and test networks.
type params struct {
	*chaincfg.Params
	DcrdRPCServerPort   string
	RPCServerPort       string
	WalletRPCServerPort string
}

// mainNetParams contains parameters specific to the main network
// (wire.MainNet).  NOTE: The RPC port is intentionally different than the
// reference implementation because wbc does not handle wallet requests.  The
// separate wallet process listens on the well-known port and forwards requests
// it does not handle on to wbc.  This approach allows the wallet process
// to emulate the full reference implementation RPC API.
var mainNetParams = params{
	Params:              &chaincfg.MainNetParams,
	DcrdRPCServerPort:   "9109",
	RPCServerPort:       "9113",
	WalletRPCServerPort: "9110",
}

// testNet2Params contains parameters specific to the test network (version 0)
// (wire.TestNet).  NOTE: The RPC port is intentionally different than the
// reference implementation - see the mainNetParams comment for details.
var testNet2Params = params{
	Params:              &chaincfg.TestNet2Params,
	DcrdRPCServerPort:   "19109",
	RPCServerPort:       "19113",
	WalletRPCServerPort: "19110",
}

// simNetParams contains parameters specific to the simulation test network
// (wire.SimNet).
var simNetParams = params{
	Params:              &chaincfg.SimNetParams,
	DcrdRPCServerPort:   "19556",
	RPCServerPort:       "19560",
	WalletRPCServerPort: "19557",
}

// netName returns the name used when referring to a wbcoin network.  At the
// time of writing, wbc currently places blocks for testnet version 0 in the
// data and log directory "testnet", which does not match the Name field of the
// chaincfg parameters.  This function can be used to override this directory name
// as "testnet" when the passed active network matches wire.TestNet.
//
// A proper upgrade to move the data and log directories for this network to
// "testnet" is planned for the future, at which point this function can be
// removed and the network parameter's name used instead.
func netName(chainParams *params) string {
	switch chainParams.Net {
	case wire.TestNet2:
		return "testnet2"
	default:
		return chainParams.Name
	}
}
