// The WBC developers. Copyright (c) 2017 
//

package wallet

import (
	"github.com/wbcoin/wbcwallet/wallet/udb"
	"github.com/wbcoin/wbcwallet/walletdb"
	"bitbucket.org/siegfriedvmblockchain/siegfried/wbc/hdkeychain"
)

// UpgradeToSLIP0044CoinType upgrades the wallet from the legacy BIP0044 coin
// type to one of the coin types assigned to WBC in SLIP0044.  This should be
// called after a new wallet is created with a random (not imported) seed.
//
// This function does not register addresses from the new account 0 with the
// wallet's network backend.  This is intentional as it allows offline
// activities, such as wallet creation, to perform this upgrade.
func (w *Wallet) UpgradeToSLIP0044CoinType() error {
	var extBranchXpub, intBranchXpub *hdkeychain.ExtendedKey

	err := walletdb.Update(w.db, func(dbtx walletdb.ReadWriteTx) error {
		err := w.Manager.UpgradeToSLIP0044CoinType(dbtx)
		if err != nil {
			return err
		}

		extBranchXpub, err = w.Manager.AccountBranchExtendedPubKey(dbtx, 0,
			udb.ExternalBranch)
		if err != nil {
			return err
		}
		intBranchXpub, err = w.Manager.AccountBranchExtendedPubKey(dbtx, 0,
			udb.InternalBranch)
		return err
	})
	if err != nil {
		return err
	}

	w.addressBuffersMu.Lock()
	w.addressBuffers[0] = &bip0044AccountData{
		albExternal: addressBuffer{branchXpub: extBranchXpub, lastUsed: ^uint32(0)},
		albInternal: addressBuffer{branchXpub: intBranchXpub, lastUsed: ^uint32(0)},
	}
	w.addressBuffersMu.Unlock()

	return nil
}
