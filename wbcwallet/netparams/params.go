// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2016-2017 The WBC developers
//

package netparams

import "bitbucket.org/siegfriedvmblockchain/siegfried/wbc/chaincfg"

// Params is used to group parameters for various networks such as the main
// network and test networks.
type Params struct {
	*chaincfg.Params
	JSONRPCClientPort string
	JSONRPCServerPort string
	GRPCServerPort    string
}

// MainNetParams contains parameters specific running dcrwallet and
// wbc on the main network (wire.MainNet).
var MainNetParams = Params{
	Params:            &chaincfg.MainNetParams,
	JSONRPCClientPort: "9109",
	JSONRPCServerPort: "9110",
	GRPCServerPort:    "9111",
}

// TestNet2Params contains parameters specific running dcrwallet and
// wbc on the test network (version 2) (wire.TestNet2).
var TestNet2Params = Params{
	Params:            &chaincfg.TestNet2Params,
	JSONRPCClientPort: "19109",
	JSONRPCServerPort: "19110",
	GRPCServerPort:    "19111",
}

// SimNetParams contains parameters specific to the simulation test network
// (wire.SimNet).
var SimNetParams = Params{
	Params:            &chaincfg.SimNetParams,
	JSONRPCClientPort: "19556",
	JSONRPCServerPort: "19557",
	GRPCServerPort:    "19558",
}
