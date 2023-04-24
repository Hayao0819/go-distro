package ostype

type F interface {
	ID() string // OSのID (空白を含まない小文字の文字列)
	Name() string // 表示用の文字列 (空白を含む文字列)
	Version() V // バージョン情報
}

// バージョン
type V interface {
	ID() string // バージョンID (空白を含まない小文字の文字列)
	CodeName() string // 表示用の文字列 (空白を含む文字列)
}

type other struct{}

func (o other) Name() string {
	return "other"
}

func (o other) ID()string{
	return "other"
}

func (o other) CodeName() string {
	return "other"
}
func (o other) Version() V {
	return o
}

var Other = other{}
