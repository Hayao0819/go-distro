//go:build !windows

package windows

import (
	"github.com/Hayao0819/go-distro/base"
)

// Get returns Other on non-Windows systems
func Get() base.OS {
	return Other
}
