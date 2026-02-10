# Go-Distro

A Go library for detecting operating systems, distributions, versions, and codenames.

## Installation

```bash
go get github.com/Hayao0819/go-distro
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/Hayao0819/go-distro"
)

func main() {
    os := distro.GetDetail()
    fmt.Printf("ID: %s\n", os.ID())
    fmt.Printf("Name: %s\n", os.FullName())
    fmt.Printf("Version: %s\n", os.VerID())
    fmt.Printf("Codename: %s\n", os.VerCodeName())
}
```

## Supported Operating Systems

### Linux Distributions

| Distribution | ID | Detection Method |
| --- | --- | --- |
| Arch Linux | `arch` | /etc/arch-release, pacman |
| Manjaro | `manjaro` | os-release ID, pacman/pamac |
| Debian | `debian` | os-release ID, dpkg |
| Ubuntu | `ubuntu` | os-release ID, dpkg |
| Linux Mint | `linuxmint` | os-release ID, dpkg |
| Fedora | `fedora` | os-release ID, dnf/rpm |
| RHEL | `rhel` | os-release ID, rpm |
| CentOS | `centos` | os-release ID, rpm |
| openSUSE Leap | `opensuse-leap` | os-release ID, zypper |
| openSUSE Tumbleweed | `opensuse-tumbleweed` | os-release ID, zypper |
| Alpine Linux | `alpine` | /etc/alpine-release, apk |
| Gentoo | `gentoo` | /etc/gentoo-release, emerge |

### macOS

Supports all major versions from Mac OS X 10.0 (Cheetah) to macOS 26 (Tahoe):

| Version | Codename |
| --- | --- |
| 10.0 - 10.15 | Cheetah, Puma, Jaguar, Panther, Tiger, Leopard, Snow Leopard, Lion, Mountain Lion, Mavericks, Yosemite, El Capitan, Sierra, High Sierra, Mojave, Catalina |
| 11 - 15 | Big Sur, Monterey, Ventura, Sonoma, Sequoia |
| 26 | Tahoe |

### Windows

| Version | Codename |
| --- | --- |
| Windows XP | Whistler |
| Windows Vista | Longhorn |
| Windows 7 | Vienna |
| Windows 8/8.1 | Blue |
| Windows 10 | Threshold |
| Windows 11 | Sun Valley |

### BSD

| OS | Detection Method |
| --- | --- |
| FreeBSD | uname -r |
| OpenBSD | uname -r |
| NetBSD | uname -r |

## Development

### Prerequisites

- Go 1.22+
- Docker (for testing)
- Nix (optional, for development environment)

### Setup with Nix

```bash
nix develop
```

### Build & Test

```bash
# Build
go build ./...

# Run tests
go test ./...

# Run Docker-based tests
make test-docker
```

## Sample Tool

Check your local environment:

```bash
go run ./cmd/main.go
```

## Contributing

Contributions are welcome! The library has a generic structure that makes it easy to add support for new distributions.

## License

You can choose whichever you prefer:

- [MIT LICENSE](./LICENSE.txt)
- [MIT-SUSHI](./SUSHI.md)

## Special Thanks

- [watasuke102/mit-sushi-ware](https://github.com/watasuke102/mit-sushi-ware)
- [dylanaraps/neofetch](https://github.com/dylanaraps/neofetch)
- [python-distro/distro](https://github.com/python-distro/distro)
