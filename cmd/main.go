package main

import (
	"strings"

	"github.com/Hayao0819/go-distro"
	"github.com/Hayao0819/go-distro/darwin"
	"github.com/Hayao0819/go-distro/goos"
	"github.com/Hayao0819/go-distro/linux"
)

func main() {
	d := distro.Get()
	if goos.Get() ==  goos.Linux{
		println("OSRelease ID      : " + linux.OSRelease.ID)
	}
	println("Detected Name     : " + d.Name())
	println("Detected ID       : " + d.ID())
	println("Detected CodeName : " + d.Version().CodeName())
	println("Supported Linux   : " + func() string {
		var s []string
		for _, l := range linux.DistroList {
			if l != linux.Other {
				s = append(s, l.ID())
			}
		}
		return strings.Join(s, ", ")
	}())
	println("Supported macOS   : " + func() string {
		var s []string
		for _, m := range darwin.VersionList {
			if m != &darwin.Other {
				s = append(s, m.Version().FullName())
			}
		}
		return strings.Join(s, ", ")
	}())
}
