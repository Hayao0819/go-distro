package linux

import (
	"github.com/Hayao0819/go-distro/base"
	"github.com/Hayao0819/go-distro/pkgmgr"
)

var Fedora = &Linux{
	id:   "fedora",
	name: "Fedora Linux",
	require: func() bool {
		if OSRelease.ID != "fedora" {
			return false
		}
		if !pkgmgr.Dnf.Installed() && !pkgmgr.Rpm.Installed() {
			return false
		}
		return true
	},
	verfunc: func() version {
		return version{
			id:       base.ID(OSRelease.VERSION_ID),
			codename: OSRelease.VERSION_CODENAME,
			fullname: OSRelease.VERSION,
		}
	},
}
