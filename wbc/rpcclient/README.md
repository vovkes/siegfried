rpcclient
=========

[![Build Status](http://img.shields.io/travis/wbcoin/wbc.svg)](https://travis-ci.org/wbcoin/wbc)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/wbcoin/wbc/rpcclient)

rpcclient implements a Websocket-enabled WBC JSON-RPC client package written
in [Go](http://golang.org/).  It provides a robust and easy to use client for
interfacing with a WBC RPC server that uses a wbc compatible WBC
JSON-RPC API.

## Status

This package is currently under active development.  It is already stable and
the infrastructure is complete.  However, there are still several RPCs left to
implement and the API is not stable yet.

## Documentation

* [API Reference](http://godoc.org/github.com/wbcoin/wbc/rpcclient)
* [wbc Websockets Example](https://github.com/wbcoin/wbc/tree/master/rpcclient/examples/wbcwebsockets)
  Connects to a wbc RPC server using TLS-secured websockets, registers for
  block connected and block disconnected notifications, and gets the current
  block count
* [dcrwallet Websockets Example](https://github.com/wbcoin/wbc/tree/master/rpcclient/examples/dcrwalletwebsockets)
  Connects to a dcrwallet RPC server using TLS-secured websockets, registers for
  notifications about changes to account balances, and gets a list of unspent
  transaction outputs (utxos) the wallet can sign

## Major Features

* Supports Websockets (wbc/dcrwallet) and HTTP POST mode (bitcoin core-like)
* Provides callback and registration functions for wbc/dcrwallet notifications
* Supports wbc extensions
* Translates to and from higher-level and easier to use Go types
* Offers a synchronous (blocking) and asynchronous API
* When running in Websockets mode (the default):
  * Automatic reconnect handling (can be disabled)
  * Outstanding commands are automatically reissued
  * Registered notifications are automatically reregistered
  * Back-off support on reconnect attempts

## Installation

```bash
$ go get -u github.com/wbcoin/wbc/rpcclient
```

## License

Package rpcclient is licensed under the [copyfree](http://copyfree.org) ISC
License.
