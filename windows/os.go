package windows

import (
	"github.com/Hayao0819/go-distro/ostype"
	//"golang.org/x/sys/windows/registry"
)

type Version struct {
	codename string
	fullname string
	id       string
}

type Win struct {
	version ostype.V
}

func (w Win) ID() string {
	return "windows"
}

func (w Win) Name() string {
	return "Windows"
}
func (w Win) Version() ostype.V {
	return w.version
}

func (v Version) ID() string {
	return v.id
}
func (v Version) CodeName() string {
	return v.codename
}

func (v Version) FullName() string {
	return v.codename
}

var (
	Other = &Win{
		version: Version{
			id:       "other",
			fullname: "other",
			codename: "other",
		},
	}
	Win7 = &Win{
		version: Version{
			id:       "7",
			codename: "vienna",
			fullname: "Vienna",
		},
	}
	Win10 = &Win{
		version: Version{
			id:       "10",
			codename: "redstone",
			fullname: "Redstone",
		},
	}
	Win11 = &Win{
		version: Version{
			id:       "11",
			fullname: "Sun Valley",
			codename: "sunvalley",
		},
	}
)

// Todo: Get Windows Version
func Get() *Win {
	return Other
}
