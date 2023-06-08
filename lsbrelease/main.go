package lsbrelease

import (
	"errors"
	"os"
	"strings"
)

// /etc/lsb-releaseの内容
type LsbRelease struct {
	ID          string
	DESCRIPTION string
	RELEASE     string
	CODENAME    string
}

// lsb-releaseのファイルパス
const Path = "/etc/lsb-release"

var ErrEmpty error = errors.New("Empty lsb-release")

func formatValue(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"")
	return s
}

func Read() (*LsbRelease, error) {
	b, err := os.ReadFile(Path)
	if err != nil {
		return nil, err
	}

	lsb := LsbRelease{}
	setcount := 0
	for _, linestr := range strings.Split(string(b), "\n") {
		line := strings.Split(linestr, "=")
		if len(line) < 2 {
			continue
		}
		value := formatValue(line[1])
		setcount++
		switch strings.TrimSpace(line[0]) {
		case "DISTRIB_ID":
			lsb.ID = value
		case "DISTRIB_DESCRIPTION":
			lsb.DESCRIPTION = value
		case "DISTRIB_RELEASE":
			lsb.RELEASE = value
		case "DISTRIB_CODENAME":
			lsb.CODENAME = value
		default:
			setcount--
		}
	}

	if setcount <= 0 {
		//return nil, nil
		// 何も代入されていない場合の処理を書く
		return nil, ErrEmpty
	}

	return &lsb, nil
}
