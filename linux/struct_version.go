package linux

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
