package linux

import (
	"os"

	"github.com/Hayao0819/go-distro/ostype"
	"github.com/Hayao0819/go-distro/pkgmgr"
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
	id: "other",
	name: "other",
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

var	Arch  = &Linux{
	id: "arch",
	name: "Arch Linux",
	verfunc: func()(ostype.V){
		return Version{
			id: "rolling",
			codename: "rolling",
		}
	},
	require: func()(bool){
		// pacmanが存在するか
		if ! pkgmgr.Pacman.Installed(){
			return false
		}
		
		// /etc/arch-releaseが存在するか
		if _, err := os.Stat("/etc/arch-release"); err != nil {
			return true
		}



		return false
	},
}

var	Debian = &Linux{id: "debian"}
var	Ubuntu = &Linux{id: "ubuntu"}
var	RHEL   = &Linux{id: "rhel"}
var	CentOS = &Linux{id: "centos"}
