package openbsd

import (
	"os/exec"
	"strings"

	"github.com/Hayao0819/go-distro/base"
)

type openBsd struct {
	version string
}

func (o openBsd) ID() base.ID {
	return "openbsd"
}

func (o openBsd) FullName() string {
	return "OpenBSD"
}

func (o openBsd) VerID() base.ID {
	return base.ID(o.version)
}

func (o openBsd) VerFullName() string {
	return "OpenBSD " + o.version
}

func (o openBsd) VerCodeName() string {
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
	return &openBsd{
		version: getVersion(),
	}
}
