package darwin

import (
	//"fmt"
	"os"

	"github.com/Hayao0819/go-distro/ostype"
	"howett.net/plist"
)

// macOSの詳細を返します
func Get() ostype.F {
	info := struct {
		BuildID                   string `plist:"BuildID"`
		ProductBuildVersion       string `plist:"ProductBuildVersion"`
		ProductCopyright          string `plist:"ProductCopyright"`
		ProductName               string `plist:"ProductName"`
		ProductUserVisibleVersion string `plist:"ProductUserVisibleVersion"`
		ProductVersion            string `plist:"ProductVersion"`
		IOSSupportVersion         string `plist:"iOSSupportVersion"`
	}{}

	sysxml, err := os.Open("/System/Library/CoreServices/SystemVersion.plist")
	if err != nil {
		return Other
	}
	decoder := plist.NewDecoder(sysxml)
	if err := decoder.Decode(&info); err != nil {
		return Other
	}
	//fmt.Println(sysxml)
	//println(info.ProductVersion)


	println(getFromVersion("10.16").Version().ID())
	return getFromVersion(info.ProductVersion)
}
