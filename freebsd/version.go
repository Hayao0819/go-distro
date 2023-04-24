package freebsd

//import "github.com/Hayao0819/go-distro/ostype"



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

