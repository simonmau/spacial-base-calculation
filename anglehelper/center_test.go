package anglehelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenter(t *testing.T) {
	testData := []struct {
		oldAngle float64
		newAngle float64
		expected float64
	}{
		{
			oldAngle: 0.0,
			newAngle: 0.0,
			expected: 0.0,
		},
		{
			oldAngle: 0.0,
			newAngle: 3.1415,
			expected: 1.57,
		},
		{
			oldAngle: 0.0,
			newAngle: 3.15,
			expected: 4.71,
		},
		{
			oldAngle: 0.0,
			newAngle: 4.0,
			expected: 5.14,
		},
	}

	for _, tc := range testData {
		value := Center(tc.oldAngle, tc.newAngle)
		assert.InDelta(t, tc.expected, value, 0.01)
	}
}
