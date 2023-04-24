package main

import (
	"strings"

	"github.com/Hayao0819/go-distro"
	"github.com/Hayao0819/go-distro/linux"
	"github.com/Hayao0819/go-distro/darwin"
)

func main(){
	d := distro.Get()
	println("Current ID      : " + d.String())
	println("Current CodeName: " + d.Version().CodeName())
	println("Supported Linux : " + func()(string){
		var s []string
		for _, l := range linux.DistroList {
			if l != linux.Other {
				s = append(s, l.String())
			}
		}
		return strings.Join(s, ", ")
	}())
	println("Supported macOS : " + func()(string){
		var s []string
		for _, m := range darwin.VersionList {
			if m != &darwin.Other {
				s = append(s, m.String())
			}
		}
		return strings.Join(s, ", ")
	}());
}
