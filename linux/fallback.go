package linux

import (
	"github.com/Hayao0819/go-distro/ostype"
	"strings"
)

// idと定数の関連付け
func getFromID(id string) *Linux {
	for _, d := range DistroList {
		if d.id == id {
			return d
		}
	}
	return Other
}

// /etc/os-releaseのIDの値から取得
// requireが定義されていないディストリのFallback
func getFromOSRelease() (ostype.F, error) {

	l := &Linux{}

	// /etc/os-releaseのIDの値から取得
	l = getFromID(OSRelease.ID)
	if l != Other {
		return l, nil
	}

	// /etc/os-releaseのID_LIKEから取得
	for _, i := range strings.Split(OSRelease.ID_LIKE, " ") {
		i = strings.TrimSpace(i)
		l = getFromID(i)
		if l != Other {
			return l, nil
		}
	}
	return Other, nil
}

// /etc/lsb-releaseのIDの値から取得
// requireが定義されていないディストリのFallback
func getFromLsbRelease() (ostype.F, error) {
	return getFromID(LSBRelease.ID), nil
}
