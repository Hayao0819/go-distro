package linux

import (
	"os"

	"github.com/Hayao0819/go-distro/ostype"
	"github.com/Hayao0819/go-distro/pkgmgr"
)

var	Ubuntu = &Linux{
	id: "ubuntu",
	name: "Ubuntu",
	require: func()(bool){
		// /etc/debian_versionが存在しない
		if _, err := os.Stat("/etc/debian_version"); err != nil {
			return false
		}

		// os-releaseのid
		if OSRelease.ID != "ubuntu" {
			return false
		}

		if ! pkgmgr.Dpkg.Installed(){
			return false
		}

		// os-releaseのidがubuntuかつdpkgがインストールされている
		return true
	},
	verfunc: func()(ostype.V){
		switch OSRelease.VERSION_ID {
			case "18.04":
				return Version{
					id: "bionic",
					codename: "Bionic Beaver",
				}
			case "16.04":
				return Version{
					id: "xenial",
					codename: "Xenial Xerus",
				}
			default:
				return Version{
					id: "none",
					codename: "none",
				}
		}
	},
}
