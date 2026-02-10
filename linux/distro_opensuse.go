package linux

import (
	"strings"

	"github.com/Hayao0819/go-distro/base"
	"github.com/Hayao0819/go-distro/pkgmgr"
)

var OpenSUSELeap = &Linux{
	id:   "opensuse-leap",
	name: "openSUSE Leap",
	require: func() bool {
		if OSRelease.ID != "opensuse-leap" {
			return false
		}
		if !pkgmgr.Zypper.Installed() {
			return false
		}
		return true
	},
	verfunc: func() version {
		return version{
			id:       base.ID(OSRelease.VERSION_ID),
			codename: strings.ToLower(strings.ReplaceAll(OSRelease.VERSION, " ", "-")),
			fullname: OSRelease.VERSION,
		}
	},
}

var OpenSUSETumbleweed = &Linux{
	id:   "opensuse-tumbleweed",
	name: "openSUSE Tumbleweed",
	require: func() bool {
		if OSRelease.ID != "opensuse-tumbleweed" {
			return false
		}
		if !pkgmgr.Zypper.Installed() {
			return false
		}
		return true
	},
	verfunc: func() version {
		return version{
			id:       "rolling",
			codename: "tumbleweed",
			fullname: "Rolling Release",
		}
	},
}
