package mathext

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func GetPointer[T Number](input T) *T {
	return &input
}
