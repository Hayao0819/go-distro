//go:build !darwin

package darwin

import "github.com/Hayao0819/go-distro/ostype"

func Get() ostype.F {
	return getFromVersion("")
}
