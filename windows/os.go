package windows

import (
	"github.com/Hayao0819/go-distro/base"
	//"golang.org/x/sys/windows/registry"
)

type win struct {
	codename string
	fullname string
	id       base.ID
}

func (w win) ID() base.ID {
	return "windows"
}

func (w win) FullName() string {
	return "Windows"
}

func (w win) VerID() base.ID {
	return base.ID(w.id)
}
func (w win) VerCodeName() string {
	return w.codename
}

func (w win) VerFullName() string {
	return w.fullname
}

var (
	Other = &win{

		id:       "other",
		fullname: "other",
		codename: "other",
	}
	Win7 = &win{

		id:       "7",
		codename: "vienna",
		fullname: "Vienna",
	}
	Win10 = &win{

		id:       "10",
		codename: "redstone",
		fullname: "Redstone",
	}
	Win11 = &win{
		id:       "11",
		fullname: "Sun Valley",
		codename: "sunvalley",
	}
)

// Todo: Get Windows Version
func Get() base.OS {
	return Other
}
