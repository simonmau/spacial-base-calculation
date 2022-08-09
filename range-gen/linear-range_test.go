package rangegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeGenArray_1(t *testing.T) {
	t.Parallel()

	min := float64(1)
	max := float64(10)
	steps := float64(9)

	data := GenLinearRangeArray(&min, &max, &steps)

	assert.Equal(t, 10, len(data))

	for i := 1; i <= 10; i++ {
		assert.Equal(t, float64(i), data[i-1])
	}
}

func TestRangeGenArray_2(t *testing.T) {
	t.Parallel()

	min := float64(1)
	max := float64(10)
	steps := float64(18)

	data := GenLinearRangeArray(&min, &max, &steps)

	assert.Equal(t, 19, len(data))
}
