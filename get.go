package distro

import (
	"github.com/Hayao0819/go-distro/goos"
	"github.com/Hayao0819/go-distro/linux"
	"github.com/Hayao0819/go-distro/darwin"
	"github.com/Hayao0819/go-distro/ostype"
)

func Get()(ostype.F){
	g := goos.Get()
	if g == goos.Linux{
		return linux.Get()
	}else if g == goos.Darwin{
		return darwin.Get()
	}else{
		return ostype.Other
	}
}
