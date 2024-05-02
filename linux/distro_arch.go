package linux

import (
	"os"

	"github.com/Hayao0819/go-distro/pkgmgr"
)

var Arch = &Linux{
	id:   "arch",
	name: "Arch Linux",
	verfunc: func() version {
		return version{
			id:       "rolling",
			fullname: "Rolling Release",
			codename: "rolling",
		}
	},
	require: func() bool {
		// /etc/arch-releaseが存在するか
		if _, err := os.Stat("/etc/arch-release"); err == nil {
			return true
		}

		// pacmanが存在するか
		if !pkgmgr.Pacman.Installed() {
			return false
		}

		// os-releaseのid
		if OSRelease.ID != "arch" {
			return false
		}

		// pacmanがインストールされている かつ os-releaseのidがarch
		return true
	},
}
