// Copyright (c) 2013-2015 The btcsuite developers
//

package legacyrpc

// Options contains the required options for running the legacy RPC server.
type Options struct {
	Username string
	Password string

	MaxPOSTClients      int64
	MaxWebsocketClients int64
}
