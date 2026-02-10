package linux

import (
	"os"

	"github.com/Hayao0819/go-distro/base"
	"github.com/Hayao0819/go-distro/pkgmgr"
)

var Gentoo = &Linux{
	id:   "gentoo",
	name: "Gentoo Linux",
	require: func() bool {
		// /etc/gentoo-release が存在するか
		if _, err := os.Stat("/etc/gentoo-release"); err == nil {
			return true
		}

		if OSRelease.ID != "gentoo" {
			return false
		}

		if !pkgmgr.Emerge.Installed() {
			return false
		}

		return true
	},
	verfunc: func() version {
		// Gentoo は VERSION_ID を持たないことが多い
		verId := OSRelease.VERSION_ID
		if verId == "" {
			verId = "rolling"
		}
		return version{
			id:       base.ID(verId),
			codename: "gentoo",
			fullname: OSRelease.PRETTY_NAME,
		}
	},
}
