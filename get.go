package distro

import (
	"github.com/Hayao0819/go-distro/darwin"
	"github.com/Hayao0819/go-distro/freebsd"
	"github.com/Hayao0819/go-distro/goos"
	"github.com/Hayao0819/go-distro/linux"
	"github.com/Hayao0819/go-distro/ostype"
	"github.com/Hayao0819/go-distro/windows"
)

// Support list
var LinuxList = &linux.DistroList
var DarwinList = &darwin.VersionList

//var WindowsList = windows.VersionList

func Get() ostype.F {
	g := goos.Get()
	switch g {
	case goos.Linux:
		return linux.Get()
	case goos.Darwin:
		return darwin.Get()
	case goos.Windows:
		return windows.Get()
	case goos.Freebsd:
		return freebsd.Get()
	default:
		return ostype.Other
	}
}
