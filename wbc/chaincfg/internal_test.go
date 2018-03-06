// The WBC developers. Copyright (c) 2017
//

package chaincfg

import (
	"testing"

	"bitbucket.org/siegfriedvmblockchain/siegfried/wbc/chaincfg/chainhash"
)

func TestInvalidHashStr(t *testing.T) {
	_, err := chainhash.NewHashFromStr("banana")
	if err == nil {
		t.Error("Invalid string should fail.")
	}
}
