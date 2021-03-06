package main

import (
	"encoding/json"

	"bitbucket.org/siegfriedvmblockchain/siegfried/wbc/chaincfg/chainhash"
	"bitbucket.org/siegfriedvmblockchain/siegfried/wbc/rpcclient"
)

// Define notification handlers
func getNodeNtfnHandlers(ctx *appContext, connCfg *rpcclient.ConnConfig) *rpcclient.NotificationHandlers {
	return &rpcclient.NotificationHandlers{
		OnNewTickets: func(blockHash *chainhash.Hash, blockHeight int64, stakeDifficulty int64, tickets []*chainhash.Hash) {
			nt := NewTicketsForBlock{
				blockHash:   blockHash,
				blockHeight: blockHeight,
				newTickets:  tickets,
			}
			ctx.newTicketsChan <- nt
		},
		OnSpentAndMissedTickets: func(blockHash *chainhash.Hash, blockHeight int64, stakeDifficulty int64, tickets map[chainhash.Hash]bool) {
			ticketsFixed := make(map[*chainhash.Hash]bool)
			for ticketHash, spent := range tickets {
				ticketHash := ticketHash
				ticketsFixed[&ticketHash] = spent
			}
			smt := SpentMissedTicketsForBlock{
				blockHash:   blockHash,
				blockHeight: blockHeight,
				smTickets:   ticketsFixed,
			}
			ctx.spentmissedTicketsChan <- smt
		},
		OnWinningTickets: func(blockHash *chainhash.Hash, blockHeight int64, winningTickets []*chainhash.Hash) {
			wt := WinningTicketsForBlock{
				blockHash:      blockHash,
				blockHeight:    blockHeight,
				winningTickets: winningTickets,
			}
			ctx.winningTicketsChan <- wt
		},
	}
}

func getWalletNtfnHandlers(cfg *config) *rpcclient.NotificationHandlers {
	return &rpcclient.NotificationHandlers{
		OnUnknownNotification: func(method string, params []json.RawMessage) {
			log.Infof("ignoring notification %v", method)
		},
	}
}
