package linux

import "github.com/Hayao0819/go-distro/ostype"

func Get() ostype.F {
	// 各ディストリビューションの判定
	for _, d := range DistroList {
		if d==Other{
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
