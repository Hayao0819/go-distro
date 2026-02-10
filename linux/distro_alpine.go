package linux

import (
	"os"

	"github.com/Hayao0819/go-distro/base"
	"github.com/Hayao0819/go-distro/pkgmgr"
)

var Alpine = &Linux{
	id:   "alpine",
	name: "Alpine Linux",
	require: func() bool {
		// /etc/alpine-release が存在するか
		if _, err := os.Stat("/etc/alpine-release"); err == nil {
			return true
		}

		if OSRelease.ID != "alpine" {
			return false
		}

		if !pkgmgr.Apk.Installed() {
			return false
		}

		return true
	},
	verfunc: func() version {
		return version{
			id:       base.ID(OSRelease.VERSION_ID),
			codename: "",
			fullname: OSRelease.PRETTY_NAME,
		}
	},
}
