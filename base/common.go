package base

// OSの各情報にアクセスできるインターフェース

type ID string

type OS interface {
	ID() ID              // OSのID (空白を含まない小文字の文字列)
	FullName() string    // 表示用の文字列 (空白を含む文字列)
	VerID() ID           // バージョンID (通常は数字の文字列)
	VerFullName() string // 表示用の文字列 (空白を含む文字列)
	VerCodeName() string // コードネーム (空白を含まない小文字の文字列)
}

type other struct{}

func (o other) ID() ID {
	return "other"
}

func (o other) FullName() string {
	return "other"
}

func (o other) VerID() ID {
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
