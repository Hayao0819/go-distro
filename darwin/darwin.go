package darwin

import (
	"strings"

	"github.com/Hayao0819/go-distro/ostype"
	"github.com/Hayao0819/sysinfo-ng"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Darwin struct {
	value   string
	version string
}

func (d Darwin) ID() string {
	return d.value
}

func (d Darwin) CodeName() string {
	// replace "hoge-fugo" to "Hoge Fugo"
	//return strings.Title(strings.ReplaceAll(d.value, "-", " "))
	return cases.Title(language.Und).String(strings.ReplaceAll(d.value, "-", " "))
}

func (d Darwin) String() string {
	return d.value
}

func (d Darwin) Version() ostype.V {
	return d
}

var (
	Other        = Darwin{value: "other"}
	Cheetah      = Darwin{value: "cheetah", version: "10.0"}
	Puma         = Darwin{value: "puma", version: "10.1"}
	Jaguar       = Darwin{value: "jaguar", version: "10.2"}
	Panther      = Darwin{value: "panther", version: "10.3"}
	Tiger        = Darwin{value: "tiger", version: "10.4"}
	Leopard      = Darwin{value: "leopard", version: "10.5"}
	ShowLeopard  = Darwin{value: "show-leopard", version: "10.6"}
	Lion         = Darwin{value: "lion", version: "10.7"}
	MountainLion = Darwin{value: "mountain-lion", version: "10.8"}
	Mavericks    = Darwin{value: "mavericks", version: "10.9"}
	Yosemite     = Darwin{value: "yosemite", version: "10.10"}
	ElCapitan    = Darwin{value: "el-capitan", version: "10.11"}
	Serra        = Darwin{value: "serra", version: "10.12"}
	HighSerra    = Darwin{value: "high-serra", version: "10.13"}
	Mojave       = Darwin{value: "mojave", version: "10.14"}
	Catalina     = Darwin{value: "catalina", version: "10.15"}
	BigSur       = Darwin{value: "big-sur", version: "11"}
	Monterey     = Darwin{value: "monterey", version: "12"}
	Ventura      = Darwin{value: "ventura", version: "13"}
)

var VersionList []*Darwin = []*Darwin{
	&Other,
	&Cheetah,
	&Puma,
	&Jaguar,
	&Panther,
	&Tiger,
	&Leopard,
	&ShowLeopard,
	&Lion,
	&MountainLion,
	&Mavericks,
	&Yosemite,
	&ElCapitan,
	&Serra,
	&HighSerra,
	&Mojave,
	&Catalina,
	&BigSur,
	&Monterey,
	&Ventura,
}

func getFromVersion(v string) ostype.F {
	for _, i := range VersionList {
		if i == &Other {
			continue
		}
		if strings.HasPrefix(v, i.Version().ID()) {
			return i
		}
	}
	return Other
}

func Get() ostype.F {
	var v sysinfo.SysInfo
	v.GetSysInfo()
	return getFromVersion(v.OS.Release)
}
