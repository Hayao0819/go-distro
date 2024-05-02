package linux

// Linuxディストリビューションを判別するための構造体
type Linux struct {
	id      string         // /etc/os-releaseのid
	name    string         // 小文字+大文字 空白あり
	verfunc func() version // バージョン情報を取得する関数
	require func() bool    // このディストリビューションであるかを判定する関数
}

type version struct {
	id       string
	codename string
	fullname string
}

func (l Linux) ID() string {
	return l.id
}

func (l Linux) FullName() string {
	return l.name
}

func (l Linux) VerID() string {
	if l.verfunc == nil {
		return "none"
	}
	return l.verfunc().id
}

func (l Linux) VerFullName() string {
	if l.verfunc == nil {
		return "none"
	}
	return l.verfunc().fullname
}

func (l Linux) VerCodeName() string {
	if l.verfunc == nil {
		return "none"
	}
	return l.verfunc().codename
}

func (l Linux) Require() bool {
	if l.require == nil {
		return false
	}
	return l.require()
}
