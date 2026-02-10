package freebsd

import (
	"os/exec"
	"strings"

	"github.com/Hayao0819/go-distro/base"
)

type freeBsd struct {
	version string
}

func (f freeBsd) ID() base.ID {
	return "freebsd"
}

func (f freeBsd) FullName() string {
	return "FreeBSD"
}

func (f freeBsd) VerID() base.ID {
	return base.ID(f.version)
}

func (f freeBsd) VerFullName() string {
	return "FreeBSD " + f.version
}

func (f freeBsd) VerCodeName() string {
	return ""
}

func getVersion() string {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func Get() base.OS {
	return &freeBsd{
		version: getVersion(),
	}
}
