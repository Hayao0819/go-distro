package goos

import (
	"runtime"
)

// 各GOOSのID
type G int

const (
	Aix G = iota
	Android
	Darwin
	Dragonfly
	Freebsd
	Illumos
	Ios
	Js
	Linux
	Netbsd
	Openbsd
	Plan9
	Solaris
	Windows
)

// runtime.GOOSの文字列から対応するIDを返します
func Get() G {
	switch runtime.GOOS {
	case "AIX":
		return Aix
	case "android":
		return Android
	case "darwin":
		return Darwin
	case "dragonfly":
		return Dragonfly
	case "freebsd":
		return Freebsd
	case "illumos":
		return Illumos
	case "ios":
		return Ios
	case "js":
		return Js
	case "linux":
		return Linux
	case "netbsd":
		return Netbsd
	case "openbsd":
		return Openbsd
	case "plan9":
		return Plan9
	case "solaris":
		return Solaris
	case "windows":
		return Windows
	default:
		return -1
	}
}

// UnixLikeなOSの一覧
var UnixLike = []G{
	Darwin,
	Dragonfly,
	Freebsd,
	Linux,
}

// UnixLikeかどうかを確認します
func (g G) IsUnix() bool {
	for _, v := range UnixLike {
		if g == v {
			return true
		}
	}
	return false
}

// 現在のOSがUnixLikeかどうか確認します
func IsUnix()bool{
	return Get().IsUnix()
}
