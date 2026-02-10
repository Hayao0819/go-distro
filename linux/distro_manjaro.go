package linux

import (
	"github.com/Hayao0819/go-distro/pkgmgr"
)

var Manjaro = &Linux{
	id:   "manjaro",
	name: "Manjaro Linux",
	require: func() bool {
		if OSRelease.ID != "manjaro" {
			return false
		}

		// pacman または pamac が存在するか
		if !pkgmgr.Pacman.Installed() && !pkgmgr.Pamac.Installed() {
			return false
		}

		return true
	},
	verfunc: func() version {
		return version{
			id:       "rolling",
			codename: "rolling",
			fullname: "Rolling Release",
		}
	},
}
