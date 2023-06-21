package darwin
import (
	"strings"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"github.com/Hayao0819/go-distro/ostype"
)

type Version struct {
	value   string
	version string
}

type OS Version

func (d Version) ID() string {
	return d.version
}

func (d Version) FullName() string {
	// replace "hoge-fugo" to "Hoge Fugo"
	//return strings.Title(strings.ReplaceAll(d.value, "-", " "))
	return cases.Title(language.Und).String(strings.ReplaceAll(d.value, "-", " "))
}

func (d Version) CodeName() string {
	return d.value
}

func (o OS) Name() string {
	return "macOS"
}

func (o OS) ID() string {
	return "darwin"
}

func (o OS) Version() ostype.V {
	return Version(o)
}
