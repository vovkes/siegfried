// Copyright (c) 2014 Conformal Systems LLC.
// The WBC developers. Copyright (c) 2017 
//

package ticketdb_test

import (
	"testing"

	"github.com/wbcoin/wbc/blockchain/stake/internal/ticketdb"
)

// TestErrorCodeStringer tests the stringized output for the ErrorCode type.
func TestErrorCodeStringer(t *testing.T) {
	tests := []struct {
		in   ticketdb.ErrorCode
		want string
	}{
		{ticketdb.ErrUndoDataShortRead, "ErrUndoDataShortRead"},
		{ticketdb.ErrUndoDataCorrupt, "ErrUndoDataCorrupt"},
		{ticketdb.ErrTicketHashesShortRead, "ErrTicketHashesShortRead"},
		{ticketdb.ErrTicketHashesCorrupt, "ErrTicketHashesCorrupt"},
		{ticketdb.ErrUninitializedBucket, "ErrUninitializedBucket"},
		{ticketdb.ErrMissingKey, "ErrMissingKey"},
		{ticketdb.ErrChainStateShortRead, "ErrChainStateShortRead"},
		{ticketdb.ErrDatabaseInfoShortRead, "ErrDatabaseInfoShortRead"},
		{ticketdb.ErrLoadAllTickets, "ErrLoadAllTickets"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestRuleError tests the error output for the RuleError type.
func TestRuleError(t *testing.T) {
	tests := []struct {
		in   ticketdb.DBError
		want string
	}{
		{ticketdb.DBError{Description: "duplicate block"},
			"duplicate block",
		},
		{ticketdb.DBError{Description: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}
