package freebsd

import "github.com/Hayao0819/go-distro/ostype"

type Bsd struct {
	id      string
	name    string
	verfunc func() ostype.V
	require func() bool
}


func (f Bsd) ID() string {
	return f.id
}

func  (f Bsd) Name() string{
	return f.name
}

func (f Bsd) Version() ostype.V {
	if f.verfunc == nil {
		return Version{
			id:       "none",
			codename: "none",
		}
	}
	return f.verfunc()
}


var FreeBSD = &Bsd{
	id:   "freebsd",
	name: "FreeBSD",
}
