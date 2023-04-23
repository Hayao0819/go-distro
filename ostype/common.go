package ostype

type F interface {
	String() string
}

type V interface{
	ID() string
	CodeName() string
}
