package darwin

import (
	//"fmt"
	"strings"

	"github.com/Hayao0819/go-distro/ostype"
)

type alias struct{
	version string
	os    *OS
}

var versionAlias = []*alias{
	{
		version: "10.16",
		os: &BigSur,
	},
}

func getFromVersion(v string) ostype.F {
	matchs := []*OS{}
	for _, i := range VersionList {
		if i == &Other {
			continue
		}

		// 完全一致
		if v == i.ID() {
			return i
		}

		// 部分一致
		if strings.HasPrefix(v, i.version) {
			matchs = append(matchs, i)
		}
	}
	
	for _, i := range versionAlias{
		if v == i.version{
			return i.os
		}
	}

	if len(matchs) != 0 {
		return matchs[len(matchs)-1]
	}

	return Other
}
