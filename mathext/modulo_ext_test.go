package mathext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosMod(t *testing.T) {
	assert.Equal(t, 0, PosMod(0, 30))
	assert.Equal(t, 0, PosMod(0, -30))

	assert.Equal(t, 15, PosMod(15, 30))
	assert.Equal(t, 15, PosMod(15, -30))

	assert.Equal(t, 1, PosMod(31, 30))
	assert.Equal(t, 1, PosMod(31, -30))

	assert.Equal(t, 1, PosMod(-29, 30))
	assert.Equal(t, 1, PosMod(-29, -30))
}
