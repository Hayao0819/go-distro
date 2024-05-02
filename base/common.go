package base

// OSの各情報にアクセスできるインターフェース
type OS interface {
	ID() string       // OSのID (空白を含まない小文字の文字列)
	FullName() string // 表示用の文字列 (空白を含む文字列)
	//CodeName() string    // コードネーム (空白を含まない小文字の文字列)
	VerID() string       // バージョンID (通常は数字の文字列)
	VerFullName() string // 表示用の文字列 (空白を含む文字列)
	VerCodeName() string // コードネーム (空白を含まない小文字の文字列)
}

type other struct{}

func (o other) ID() string {
	return "other"
}

func (o other) FullName() string {
	return "other"
}

func (o other) VerID() string {
	return "other"
}

func (o other) VerFullName() string {
	return "other"
}

func (o other) VerCodeName() string {
	return "other"
}

// 未定義のOS
var Other = other{}
