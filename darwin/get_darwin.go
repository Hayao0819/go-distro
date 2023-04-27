//go:build darwin

package darwin

import (
	"github.com/Hayao0819/sysinfo-ng"
	"github.com/Hayao0819/go-distro/ostype"
)

func Get() ostype.F {
	var v sysinfo.SysInfo
	v.GetSysInfo()
	return getFromVersion(v.OS.Release)
}
