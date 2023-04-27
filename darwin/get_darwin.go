//go:build darwin

package darwin

import (
	"os"

	"github.com/Hayao0819/go-distro/ostype"
	"howett.net/plist"
)

func Get() ostype.F {
	info := struct{
		release string `plist:"ProductVersion"`
	}{}


	sysxml, err := os.Open("/System/Library/CoreServices/SystemVersion.plist") 
	if err !=nil{
		return Other
	}
	decoder := plist.NewDecoder(sysxml)
	if err := decoder.Decode(&info); err !=nil{
		return Other
	}

	return getFromVersion(info.release)
}
