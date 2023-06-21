package linux

import (
	"os"
	"strings"

	"github.com/Hayao0819/go-distro/ostype"
	"github.com/Hayao0819/go-distro/pkgmgr"
)

var Ubuntu = &Linux{
	id:   "ubuntu",
	name: "Ubuntu",
	require: func() bool {
		// /etc/debian_versionが存在しない
		if _, err := os.Stat("/etc/debian_version"); err != nil {
			return false
		}

		// os-releaseのid
		if OSRelease.ID != "ubuntu" {
			return false
		}

		if !pkgmgr.Dpkg.Installed() {
			return false
		}

		// os-releaseのidがubuntuかつdpkgがインストールされている
		return true
	},
	verfunc: func() ostype.V {

		codename := OSRelease.ADDITIONAL_FIELDS["UBUNTU_CODENAME"]
		if strings.TrimSpace(codename) == "" {
			codename = OSRelease.VERSION_CODENAME
		}

		switch codename {
		case "bionic":
			return Version{
				id:       "18.04",
				codename: "bionic",
				fullname: "Bionic Beaver",
			}
		case "xenial":
			return Version{
				id:       "16.04",
				codename: "xenial",
				fullname: "Xenial Xerus",
			}
		default:
			return Version{
				id:       OSRelease.VERSION_ID,
				codename: OSRelease.VERSION_CODENAME,
				fullname: strings.TrimSuffix(strings.Split(OSRelease.VERSION, "(")[1], ")"),
			}
		}
	},
}
