package ostype

// OSの各情報にアクセスできるインターフェース
type F interface {
	ID() string // OSのID (空白を含まない小文字の文字列)
	Name() string // 表示用の文字列 (空白を含む文字列)
	Version() V // バージョン情報
}

// OSのバージョン情報の詳細にアクセスできるインターフェース
type V interface {
	ID() string // バージョンID (通常は数字の文字列)
	FullName() string // 表示用の文字列 (空白を含む文字列)
	CodeName() string // コードネーム (空白を含まない小文字の文字列)
}

// interface Fを満たす、未定義のOSで使われる構造体
type other struct{}

func (o other) Name() string {
	return "other"
}

func (o other) ID()string{
	return "other"
}

func (o other) CodeName() string {
	return "Other"
}

func (o other) FullName() string {
	return "other"
}

func (o other) Version() V {
	return o
}

// 未定義のOS
var Other = other{}
