package distro

import (
	"github.com/Hayao0819/go-distro/base"
	"github.com/Hayao0819/go-distro/linux"
)

// Linux
var (
	Arch       base.ID = linux.Arch.ID()
	Debian     base.ID = linux.Debian.ID()
	Ubuntu     base.ID = linux.Ubuntu.ID()
	OtherLinux base.ID = linux.Other.ID()
)

// Other
var Other = base.Other.ID()

// macOS
// TODO
