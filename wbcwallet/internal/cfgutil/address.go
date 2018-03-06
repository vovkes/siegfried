// Copyright (c) 2015-2016 The btcsuite developers
// Copyright (c) 2016 The WBC developers
//

package cfgutil

import "bitbucket.org/siegfriedvmblockchain/siegfried/wbc/dcrutil"

// AddressFlag embeds a wbcutil.Address and implements the flags.Marshaler and
// Unmarshaler interfaces so it can be used as a config struct field.
type AddressFlag struct {
	dcrutil.Address
}

// NewAddressFlag creates an AddressFlag with a default wbcutil.Address.
func NewAddressFlag(defaultValue dcrutil.Address) *AddressFlag {
	return &AddressFlag{defaultValue}
}

// MarshalFlag satisifes the flags.Marshaler interface.
func (a *AddressFlag) MarshalFlag() (string, error) {
	if a.Address != nil {
		return a.Address.String(), nil
	}

	return "", nil
}

// UnmarshalFlag satisifes the flags.Unmarshaler interface.
func (a *AddressFlag) UnmarshalFlag(addr string) error {
	if addr == "" {
		a.Address = nil
		return nil
	}
	address, err := dcrutil.DecodeAddress(addr)
	if err != nil {
		return err
	}
	a.Address = address
	return nil
}
