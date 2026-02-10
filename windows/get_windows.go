//go:build windows

package windows

import (
	"github.com/Hayao0819/go-distro/base"
	"golang.org/x/sys/windows/registry"
)

// Get returns the Windows version information
func Get() base.OS {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return Other
	}
	defer k.Close()

	// Get CurrentMajorVersionNumber (Windows 10+)
	majorVersion, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
	if err == nil {
		// Get CurrentBuildNumber for Windows 11 detection
		buildNumber, _, _ := k.GetStringValue("CurrentBuildNumber")

		if majorVersion >= 10 {
			// Windows 11 starts from build 22000
			if buildNumber >= "22000" {
				return Win11
			}
			return Win10
		}
	}

	// Fallback to CurrentVersion for older Windows
	currentVersion, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		return Other
	}

	switch currentVersion {
	case "5.1":
		return WinXP
	case "6.0":
		return WinVista
	case "6.1":
		return Win7
	case "6.2":
		return Win8
	case "6.3":
		return Win81
	default:
		return Other
	}
}
