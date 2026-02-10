package pkgmgr

import "os/exec"

type P struct {
	value string
}

func (p *P) String() string {
	return p.value
}

func (p *P) Installed() bool {
	_, e := exec.LookPath(p.String())
	return e == nil
}

var (
	Pacman = P{"pacman"}
	Dpkg   = P{"dpkg"}
	Rpm    = P{"rpm"}
	Apk    = P{"apk"}
	Dnf    = P{"dnf"}
	Yum    = P{"yum"}
	Zypper = P{"zypper"}
	Emerge = P{"emerge"}
	Pamac  = P{"pamac"}
)

var PList = []*P{
	&Pacman,
	&Dpkg,
	&Rpm,
	&Apk,
	&Dnf,
	&Yum,
	&Zypper,
	&Emerge,
	&Pamac,
}

func LookFor() *P {
	for _, p := range PList {
		if p.Installed() {
			return p
		}
	}
	return nil
}
