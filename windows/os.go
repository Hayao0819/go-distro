package windows

import (
	"github.com/Hayao0819/go-distro/base"
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
		fullname: "Unknown",
		codename: "other",
	}
	WinXP = &win{
		id:       "5.1",
		codename: "whistler",
		fullname: "Whistler",
	}
	WinVista = &win{
		id:       "6.0",
		codename: "longhorn",
		fullname: "Longhorn",
	}
	Win7 = &win{
		id:       "7",
		codename: "vienna",
		fullname: "Vienna",
	}
	Win8 = &win{
		id:       "8",
		codename: "blue",
		fullname: "Blue",
	}
	Win81 = &win{
		id:       "8.1",
		codename: "blue",
		fullname: "Blue",
	}
	Win10 = &win{
		id:       "10",
		codename: "threshold",
		fullname: "Threshold",
	}
	Win11 = &win{
		id:       "11",
		fullname: "Sun Valley",
		codename: "sunvalley",
	}
)

var VersionList = []*win{
	WinXP,
	WinVista,
	Win7,
	Win8,
	Win81,
	Win10,
	Win11,
	Other,
}
