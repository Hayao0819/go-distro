package linux

import (
	//"os"

	//"github.com/Hayao0819/go-distro/ostype"
	"github.com/Hayao0819/go-distro/pkgmgr"
)

var Debian = &Linux{
	id:   "debian",
	name: "Debian",
	require: func() bool {
		// os-releaseのid
		if OSRelease.ID != "debian" {
			return false
		}

		if !pkgmgr.Dpkg.Installed() {
			return false
		}

		// os-releaseのidがdebianかつdpkgがインストールされている
		return true
	},
}
