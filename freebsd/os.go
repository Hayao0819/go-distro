package freebsd

import (
	"github.com/Hayao0819/go-distro/base"
)

type freeBsd struct {
	verFullName string
	verCodeName string
}

func (f freeBsd) ID() string {
	return "freebsd"
}

func (f freeBsd) FullName() string {
	return "FreeBSD"
}

func (f freeBsd) VerID() string {
	return "none"
}

func (f freeBsd) VerFullName() string {
	return f.verFullName
}

func (f freeBsd) VerCodeName() string {
	return f.verCodeName
}

var FreeBSD = &freeBsd{}

func Get() base.OS {
	return FreeBSD
}
