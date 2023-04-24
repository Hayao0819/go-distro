package linux

import (
	"github.com/Hayao0819/go-distro/lsbrelease"
	"github.com/Hayao0819/go-distro/ostype"
	"github.com/ashcrow/osrelease"
)

func Get() ostype.F {
	// os-releaseの読み込み
	OSRelease, _ = osrelease.New()
	LSBRelease, _ = lsbrelease.Read()

	// 各ディストリビューションの判定
	for _, d := range DistroList {
		if d == Other {
			continue
		}
		if d.require != nil && d.require() {
			return d
		}
	}

	// /etc/lsb-releaseのIDの値から取得
	if l, err := getFromLsbRelease(); err == nil && l != Other {
		return l
	}

	// /etc/os-releaseのIDの値から取得
	if l, err := getFromOSRelease(); err == nil && l != Other {
		return l
	}

	return Other
}
