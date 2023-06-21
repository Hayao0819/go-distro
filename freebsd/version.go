package freebsd

//import "github.com/Hayao0819/go-distro/ostype"

type Version struct {
	id       string
	fullname string
	codename string
}

func (v Version) ID() string {
	return v.id
}

func (v Version) CodeName() string {
	return v.codename
}

func (v Version) FullName() string {
	return v.fullname
}
