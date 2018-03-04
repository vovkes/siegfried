// Copyright (c) 2013-2014 The btcsuite developers
// The WBC developers. Copyright (c) 2017 
//

package limits

// SetLimits is a no-op on Plan 9 due to the lack of process accounting.
func SetLimits() error {
	return nil
}
