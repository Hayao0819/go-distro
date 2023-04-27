package darwin

import (
	"strings"

	"github.com/Hayao0819/go-distro/ostype"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)


type Version struct {
	value   string
	version string
}

type OS Version

func (d Version) ID() string {
	return d.value
}

func (d Version) FullName() string {
	// replace "hoge-fugo" to "Hoge Fugo"
	//return strings.Title(strings.ReplaceAll(d.value, "-", " "))
	return cases.Title(language.Und).String(strings.ReplaceAll(d.value, "-", " "))
}

func (d Version) CodeName() string {
	return d.version
}

func (o OS) Name() string {
	return "macOS"
}

func (o OS) ID() string {
	return "darwin"
}

func (o OS) Version() ostype.V {
	return Version(o)
}

var (
	Other        = OS{value: "other"}
	Cheetah      = OS{value: "cheetah", version: "10.0"}
	Puma         = OS{value: "puma", version: "10.1"}
	Jaguar       = OS{value: "jaguar", version: "10.2"}
	Panther      = OS{value: "panther", version: "10.3"}
	Tiger        = OS{value: "tiger", version: "10.4"}
	Leopard      = OS{value: "leopard", version: "10.5"}
	ShowLeopard  = OS{value: "show-leopard", version: "10.6"}
	Lion         = OS{value: "lion", version: "10.7"}
	MountainLion = OS{value: "mountain-lion", version: "10.8"}
	Mavericks    = OS{value: "mavericks", version: "10.9"}
	Yosemite     = OS{value: "yosemite", version: "10.10"}
	ElCapitan    = OS{value: "el-capitan", version: "10.11"}
	Serra        = OS{value: "serra", version: "10.12"}
	HighSerra    = OS{value: "high-serra", version: "10.13"}
	Mojave       = OS{value: "mojave", version: "10.14"}
	Catalina     = OS{value: "catalina", version: "10.15"}
	BigSur       = OS{value: "big-sur", version: "11"}
	Monterey     = OS{value: "monterey", version: "12"}
	Ventura      = OS{value: "ventura", version: "13"}
)

var VersionList []*OS = []*OS{
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


