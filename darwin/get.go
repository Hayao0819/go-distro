package darwin

import (
	//"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Hayao0819/go-distro/base"
	"howett.net/plist"
)

// macOSの詳細を返します
func getVersionIdFromPlist() string {
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
		return ""
	}
	decoder := plist.NewDecoder(sysxml)
	if err := decoder.Decode(&info); err != nil {
		return ""
	}
	//fmt.Println(sysxml)
	//println(info.ProductVersion)

	return info.ProductVersion
}

// sw_versの出力結果からmacOSのProductVersionを返します
func getVersionIdFromSwVers() string {
	cmdstr := "sw_vers | grep ProductVersion | cut -f 3"
	out, err := exec.Command("sh", "-c", cmdstr).Output()

	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(out))
}

func Get() base.OS {
	version := getVersionIdFromSwVers()
	if version == "" {
		version = getVersionIdFromPlist()
	}

	return getFromVersion(version)
}
