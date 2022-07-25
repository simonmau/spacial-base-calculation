package anglehelper

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModulo_1(t *testing.T) {
	assert.Equal(t, 0.0, Modulo(0.0))
}

func TestModulo_2(t *testing.T) {
	assert.Equal(t, 0.0, Modulo(math.Pi*2.0))
}

func TestModulo_3(t *testing.T) {
	assert.Equal(t, 0.0, Modulo(-math.Pi*2.0))
}

func TestModulo_4(t *testing.T) {
	//expected: math32.Pi*2.0-1.0
	assert.Equal(t, 5.283185, Modulo(-math.Pi*4.0-1.0))
}
