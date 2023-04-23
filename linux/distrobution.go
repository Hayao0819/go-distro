package linux

import (
	"os"

	"github.com/Hayao0819/go-distro/ostype"
)

var DistroList []*Linux = []*Linux{
	Other,
	Arch,
	Debian,
	Ubuntu,
	RHEL,
	CentOS,
}

var Other  = &Linux{
	value: "other",
	verfunc: func()(ostype.V){
		return Version{
			id: "none",
			codename: "none",
		}
	},
	require: func()(bool){
		return true
	},
}

var	Arch   = &Linux{
	value: "arch",
	verfunc: func()(ostype.V){
		return Version{
			id: "rolling",
			codename: "rolling",
		}
	},
	require: func()(bool){
		// /etc/arch-releaseが存在するか
		if _, err := os.Stat("/etc/arch-release"); err != nil {
			return true
		}

		return false
	},
}
var	Debian = &Linux{value: "debian"}
var	Ubuntu = &Linux{value: "ubuntu"}
var	RHEL   = &Linux{value: "rhel"}
var	CentOS = &Linux{value: "centos"}
