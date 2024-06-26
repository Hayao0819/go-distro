package main

import (
	"strings"

	"github.com/Hayao0819/go-distro"
	"github.com/Hayao0819/go-distro/darwin"
	"github.com/Hayao0819/go-distro/goos"
	"github.com/Hayao0819/go-distro/linux"
)

func main() {
	d := distro.GetDetail()
	if goos.Get() == goos.Linux {
		println("OSRelease ID      : " + linux.OSRelease.ID)
	}
	println("Detected Name     : " + d.FullName())
	println("Detected ID       : " + d.ID())
	println("Detected CodeName : " + d.VerCodeName())
	println("Supported Linux   : " + func() string {
		var s []string
		for _, l := range linux.DistroList {
			if l != linux.Other {
				s = append(s, string(l.ID()))
			}
		}
		return strings.Join(s, ", ")
	}())
	println("Supported macOS   : " + func() string {
		var s []string
		for _, m := range darwin.VersionList {
			if m != &darwin.Other {
				s = append(s, m.VerFullName())
			}
		}
		return strings.Join(s, ", ")
	}())
}
