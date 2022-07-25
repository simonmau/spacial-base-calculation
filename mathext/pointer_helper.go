package mathext

type Number interface {
	~float64 | ~int32 | ~int
}

func GetPointer[T Number](input T) *T {
	return &input
}
