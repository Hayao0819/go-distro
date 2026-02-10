package netbsd

import (
	"os/exec"
	"strings"

	"github.com/Hayao0819/go-distro/base"
)

type netBsd struct {
	version string
}

func (n netBsd) ID() base.ID {
	return "netbsd"
}

func (n netBsd) FullName() string {
	return "NetBSD"
}

func (n netBsd) VerID() base.ID {
	return base.ID(n.version)
}

func (n netBsd) VerFullName() string {
	return "NetBSD " + n.version
}

func (n netBsd) VerCodeName() string {
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
	return &netBsd{
		version: getVersion(),
	}
}
