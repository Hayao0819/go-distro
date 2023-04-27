package darwin

import (
	"os"

	"github.com/Hayao0819/go-distro/ostype"
	"howett.net/plist"
)

func Get() ostype.F {
	info := struct{
		BuildID string `plist:"BuildID"`
		ProductBuildVersion string `plist:"ProductBuildVersion"`
		ProductCopyright string `plist:"ProductCopyright"`
		ProductName string `plist:"ProductName"`
		ProductUserVisibleVersion string `plist:"ProductUserVisibleVersion"`
		ProductVersion string `plist:"ProductVersion"`
		IOSSupportVersion string `plist:"iOSSupportVersion"`
	}{}


	sysxml, err := os.Open("/System/Library/CoreServices/SystemVersion.plist")
	if err !=nil{
		return Other
	}
	decoder := plist.NewDecoder(sysxml)
	if err := decoder.Decode(&info); err !=nil{
		return Other
	}
	println(info.ProductVersion)

	return getFromVersion(info.ProductVersion)
}
