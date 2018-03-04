// Copyright (c) 2013-2014 The btcsuite developers
// The WBC developers. Copyright (c) 2017 
//

package limits

// SetLimits is a no-op on Windows since it's not required there.
func SetLimits() error {
	return nil
}
