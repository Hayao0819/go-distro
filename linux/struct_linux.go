package linux

import "github.com/Hayao0819/go-distro/ostype"

type Linux struct {
	id      string          // /etc/os-releaseのid
	name    string          // 小文字+大文字 空白あり
	verfunc func() ostype.V // バージョン情報を取得する関数
	require func() bool     // このディストリビューションであるかを判定する関数
}

func (l Linux) ID() string {
	return l.id
}

func (l Linux)Name()string{
	return l.name
}

func (l Linux) Version() ostype.V {
	if l.verfunc == nil {
		return Version{
			id:       "none",
			codename: "none",
		}
	}
	return l.verfunc()
}

func (l Linux) Require() bool {
	if l.require == nil {
		return false
	}
	return l.require()
}
