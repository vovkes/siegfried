// The WBC developers. Copyright (c) 2017
//

// This file should compiled from the commit the file was introduced, otherwise
// it may not compile due to API changes, or may not create the database with
// the correct old version.  This file should not be updated for API changes.

package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"bitbucket.org/siegfriedvmblockchain/siegfried/wbc/chaincfg"
	"bitbucket.org/siegfriedvmblockchain/siegfried/wbcutil/hdkeychain"
	"bitbucket.org/siegfriedvmblockchain/siegfried/wbcwallet/wallet/udb"
	"bitbucket.org/siegfriedvmblockchain/siegfried/wbcwallet/walletdb"
	_ "bitbucket.org/siegfriedvmblockchain/siegfried/wbcwallet/walletdb/bdb"
	"bitbucket.org/siegfriedvmblockchain/siegfried/wbcwallet/walletseed"
)

const dbname = "v3.db"

var (
	pubPass  = []byte("public")
	privPass = []byte("private")
)

var chainParams = &chaincfg.TestNet2Params

func main() {
	err := setup()
	if err != nil {
		fmt.Fprintf(os.Stderr, "setup: %v\n", err)
		os.Exit(1)
	}
	err = compress()
	if err != nil {
		fmt.Fprintf(os.Stderr, "compress: %v\n", err)
		os.Exit(1)
	}
}

func setup() error {
	db, err := walletdb.Create("bdb", dbname)
	if err != nil {
		return err
	}
	defer db.Close()
	seed, err := walletseed.GenerateRandomSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		return err
	}
	return udb.Initialize(db, chainParams, seed, pubPass, privPass)
}

func compress() error {
	db, err := os.Open(dbname)
	if err != nil {
		return err
	}
	defer os.Remove(dbname)
	defer db.Close()
	dbgz, err := os.Create(dbname + ".gz")
	if err != nil {
		return err
	}
	defer dbgz.Close()
	gz := gzip.NewWriter(dbgz)
	_, err = io.Copy(gz, db)
	if err != nil {
		return err
	}
	return gz.Close()
}
