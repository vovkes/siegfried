// Copyright (c) 2016 The WBC developers
//

package main

import "fmt"

type semver struct {
	major, minor, patch uint32
}

func semverCompatible(required, actual semver) bool {
	switch {
	case required.major != actual.major:
		return false
	case required.minor > actual.minor:
		return false
	case required.minor == actual.minor && required.patch > actual.patch:
		return false
	default:
		return true
	}
}

func (s semver) String() string {
	return fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)
}
