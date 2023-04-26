package linux

type Version struct {
	id       string // 通常は数字のみの文字列
	codename string // 空白を含まない小文字の文字列
	fullname string // 空白を含む表示用の文字列
}

func (v Version) ID() string {
	return v.id
}

func (v Version) CodeName() string {
	return v.codename
}

func (v Version) FullName() string {
	return v.fullname
}
