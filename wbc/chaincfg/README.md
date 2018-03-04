chaincfg
========

[![Build Status](http://img.shields.io/travis/wbcoin/wbc.svg)](https://travis-ci.org/wbcoin/wbc)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/wbcoin/wbc/chaincfg)

Package chaincfg defines chain configuration parameters for the three standard
WBC networks and provides the ability for callers to define their own custom
WBC networks.

Although this package was primarily written for wbc, it has intentionally been
designed so it can be used as a standalone package for any projects needing to
use parameters for the standard WBC networks or for projects needing to
define their own network.

## Sample Use

```Go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/wbcoin/wbc/dcrutil"
	"github.com/wbcoin/wbc/chaincfg"
)

var testnet = flag.Bool("testnet", false, "operate on the testnet WBC network")

// By default (without -testnet), use mainnet.
var chainParams = &chaincfg.MainNetParams

func main() {
	flag.Parse()

	// Modify active network parameters if operating on testnet.
	if *testnet {
		chainParams = &chaincfg.TestNetParams
	}

	// later...

	// Create and print new payment address, specific to the active network.
	pubKeyHash := make([]byte, 20)
	addr, err := btcutil.NewAddressPubKeyHash(pubKeyHash, chainParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr)
}
```

## Installation and Updating

```bash
$ go get -u github.com/wbcoin/wbc/chaincfg
```

## License

Package chaincfg is licensed under the [copyfree](http://copyfree.org) ISC
License.
