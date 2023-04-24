package freebsd

import "github.com/Hayao0819/go-distro/ostype"

type FreeBSD struct {
	id      string
	name    string
	verfunc func() ostype.V
	require func() bool
}

type Version struct {
	id       string
	codename string
}

func (v Version) ID() string {
	return v.id
}

func (v Version) CodeName() string {
	return v.codename
}

func (f FreeBSD) String() string {
	return f.id
}

func (f FreeBSD) Version() ostype.V {
	if f.verfunc == nil {
		return Version{
			id:       "none",
			codename: "none",
		}
	}
	return f.verfunc()
}
