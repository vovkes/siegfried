// The WBC developers. Copyright (c) 2017 
//

package util

import (
	"crypto/elliptic"
	"io/ioutil"
	"os"
	"time"

	"bitbucket.org/siegfriedvmblockchain/siegfried/wbc/certgen"
)

// GenCertPair generates a key/cert pair to the paths provided.
func GenCertPair(org, certFile, keyFile string) error {
	validUntil := time.Now().Add(10 * 365 * 24 * time.Hour)
	cert, key, err := certgen.NewTLSCertPair(elliptic.P521(), org,
		validUntil, nil)
	if err != nil {
		return err
	}

	// Write cert and key files.
	if err = ioutil.WriteFile(certFile, cert, 0666); err != nil {
		return err
	}
	if err = ioutil.WriteFile(keyFile, key, 0600); err != nil {
		os.Remove(certFile)
		return err
	}

	return nil
}
