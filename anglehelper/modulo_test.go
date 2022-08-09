package anglehelper

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModulo_1(t *testing.T) {
	assert.InDelta(t, 0.0, Modulo(0.0), 0.000001)
}

func TestModulo_2(t *testing.T) {
	assert.InDelta(t, 0.0, Modulo(math.Pi*2.0), 0.000001)
}

func TestModulo_3(t *testing.T) {
	assert.InDelta(t, 0.0, Modulo(-math.Pi*2.0), 0.000001)
}

func TestModulo_4(t *testing.T) {
	//expected: math32.Pi*2.0-1.0
	assert.InDelta(t, 5.283185, Modulo(-math.Pi*4.0-1.0), 0.000001)
}

//makes 1000 runs per operation -> divide result by 1000
func BenchmarkModulo(b *testing.B) {
	_ = Modulo(math.Pi)

	for i := 0; i < b.N; i++ {
		for y := 1.0; y < 1000.0; y++ {
			_ = Modulo(y)
		}
	}
}
