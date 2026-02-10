package linux

import (
	"github.com/Hayao0819/go-distro/base"
	"github.com/Hayao0819/go-distro/pkgmgr"
)

var Other = &Linux{
	id:   "other",
	name: "Other Linux",
	verfunc: func() version {
		return version{
			id:       "none",
			codename: "none",
			fullname: "Unknown",
		}
	},
	require: func() bool {
		return true
	},
}

var RHEL = &Linux{
	id:   "rhel",
	name: "Red Hat Enterprise Linux",
	require: func() bool {
		if OSRelease.ID != "rhel" {
			return false
		}
		if !pkgmgr.Rpm.Installed() {
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

var CentOS = &Linux{
	id:   "centos",
	name: "CentOS",
	require: func() bool {
		if OSRelease.ID != "centos" {
			return false
		}
		if !pkgmgr.Rpm.Installed() {
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
