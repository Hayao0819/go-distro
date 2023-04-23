package linux

import (
	"strings"

	"github.com/Hayao0819/go-distro/lsbrelease"
	"github.com/Hayao0819/go-distro/ostype"
	"github.com/ashcrow/osrelease"
)


func getFromID(id string) *Linux {
	switch id {
	case "debian":
		return Debian
	case "arch":
		return Arch
	case "ubuntu":
		return Ubuntu
	case "centos":
		return CentOS
	default:
		return Other
	}
}

func getFromOSRelease() (ostype.F, error) {
	o, err := osrelease.New()
	if err != nil {
		return nil, err
	}

	l := &Linux{}

	// /etc/os-releaseのIDの値から取得
	l = getFromID(o.ID)
	if l != Other {
		return l, nil
	}

	// /etc/os-releaseのID_LIKEから取得
	for _, i := range strings.Split(o.ID_LIKE, " ") {
		i = strings.TrimSpace(i)
		l = getFromID(i)
		if l != Other {
			return l, nil
		}
	}
	return Other, nil
}

func getFromLsbRelease() (ostype.F, error) {
	l, err := lsbrelease.Read()
	if err != nil {
		return nil, err
	}
	return getFromID(l.ID), nil
}
