package darwin

var (
	Other        = Version{value: "other"}
	Cheetah      = Version{value: "cheetah", version: "10.0",}
	Puma         = Version{value: "puma", version: "10.1"}
	Jaguar       = Version{value: "jaguar", version: "10.2"}
	Panther      = Version{value: "panther", version: "10.3"}
	Tiger        = Version{value: "tiger", version: "10.4"}
	Leopard      = Version{value: "leopard", version: "10.5"}
	ShowLeopard  = Version{value: "show-leopard", version: "10.6"}
	Lion         = Version{value: "lion", version: "10.7"}
	MountainLion = Version{value: "mountain-lion", version: "10.8"}
	Mavericks    = Version{value: "mavericks", version: "10.9"}
	Yosemite     = Version{value: "yosemite", version: "10.10"}
	ElCapitan    = Version{value: "el-capitan", version: "10.11"}
	Serra        = Version{value: "serra", version: "10.12"}
	HighSerra    = Version{value: "high-serra", version: "10.13"}
	Mojave       = Version{value: "mojave", version: "10.14"}
	Catalina     = Version{value: "catalina", version: "10.15"}
	BigSur       = Version{value: "big-sur", version: "11"}
	Monterey     = Version{value: "monterey", version: "12"}
	Ventura      = Version{value: "ventura", version: "13"}
)

var VersionList []*Version = []*Version{
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
