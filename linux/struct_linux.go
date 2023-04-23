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

func (l Linux) GetVersion()(ostype.V){
	return l.verfunc()
}

func (l Linux)Require()(bool){
	return l.require()
}
