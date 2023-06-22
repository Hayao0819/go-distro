package darwin

import (
	//"fmt"
	"os/exec"
	"strings"

	"github.com/Hayao0819/go-distro/ostype"
)

// sw_versの出力結果からmacOSのProductVersionを返します
func Get() ostype.F {
	cmdstr := "sw_vers | grep ProductVersion | cut -f 3"
	out, err := exec.Command("sh", "-c", cmdstr).Output()

	if err != nil {
		return Other
	}

	ProductVersion := strings.TrimSpace(string(out))

	return getFromVersion(ProductVersion)
}
