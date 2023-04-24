package ostype

type F interface {
	String() string
	Version() V
}

// バージョン
type V interface {
	ID() string
	CodeName() string
}

type other struct{}

func (o other) String() string {
	return "other"
}
func (o other) ID() string {
	return "other"
}
func (o other) CodeName() string {
	return "other"
}
func (o other) Version() V {
	return o
}

var Other = other{}
