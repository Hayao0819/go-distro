package linux

import "github.com/Hayao0819/go-distro/ostype"

type Linux struct {
	value string
	verfunc func()(ostype.V)
	require func()(bool)
}

func (l Linux) String() string {
	return l.value
}
func (l Linux)Version() ostype.V{
	if l.verfunc == nil{
		return Version{
			id: "none",
			codename: "none",
		}
	}
	return l.verfunc()
}


func (l Linux)Require()(bool){
	if l.require == nil{
		return false
	}
	return l.require()
}
