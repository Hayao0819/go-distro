package darwin

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Version struct {
	value   string
	version string
}

func (d Version) ID() string {
	return d.version
}

func (d Version) VerFullName() string {
	// replace "hoge-fugo" to "Hoge Fugo"
	//return strings.Title(strings.ReplaceAll(d.value, "-", " "))
	return cases.Title(language.Und).String(strings.ReplaceAll(d.value, "-", " "))
}

func (d Version) VerCodeName() string {
	return d.value
}

func (Version) FullName() string {
	return "macOS"
}

func (Version) VerID() string {
	return "darwin"
}
