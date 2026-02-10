package linux

import (
	"github.com/Hayao0819/go-distro/base"
	"github.com/Hayao0819/go-distro/pkgmgr"
)

var LinuxMint = &Linux{
	id:   "linuxmint",
	name: "Linux Mint",
	require: func() bool {
		if OSRelease.ID != "linuxmint" {
			return false
		}

		if !pkgmgr.Dpkg.Installed() {
			return false
		}

		return true
	},
	verfunc: func() version {
		codename := OSRelease.VERSION_CODENAME
		return version{
			id:       base.ID(OSRelease.VERSION_ID),
			codename: codename,
			fullname: OSRelease.VERSION,
		}
	},
}
