// Copyright (c) 2015-2016 The btcsuite developers
// Copyright (c) 2016 The WBC developers
//

package cfgutil

import (
	"strconv"
	"strings"

	"github.com/wbcoin/wbc/dcrutil"
)

// AmountFlag embeds a wbcutil.Amount and implements the flags.Marshaler and
// Unmarshaler interfaces so it can be used as a config struct field.
type AmountFlag struct {
	dcrutil.Amount
}

// NewAmountFlag creates an AmountFlag with a default wbcutil.Amount.
func NewAmountFlag(defaultValue dcrutil.Amount) *AmountFlag {
	return &AmountFlag{defaultValue}
}

// MarshalFlag satisifes the flags.Marshaler interface.
func (a *AmountFlag) MarshalFlag() (string, error) {
	return a.Amount.String(), nil
}

// UnmarshalFlag satisifes the flags.Unmarshaler interface.
func (a *AmountFlag) UnmarshalFlag(value string) error {
	value = strings.TrimSuffix(value, " DCR")
	valueF64, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	amount, err := dcrutil.NewAmount(valueF64)
	if err != nil {
		return err
	}
	a.Amount = amount
	return nil
}
