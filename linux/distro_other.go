package linux

import (
	//"os"

	"github.com/Hayao0819/go-distro/ostype"
	//"github.com/Hayao0819/go-distro/pkgmgr"
)



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






var	RHEL   = &Linux{id: "rhel"}
var	CentOS = &Linux{id: "centos"}
