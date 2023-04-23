package linux

import "github.com/Hayao0819/go-distro/ostype"

func Get() ostype.F {
	if l, err := getFromLsbRelease(); err == nil && l != Other {
		return l
	}

	if l, err := getFromOSRelease(); err == nil && l != Other {
		return l
	}

	return Other
}
