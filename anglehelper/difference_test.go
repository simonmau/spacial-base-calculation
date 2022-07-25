package anglehelper

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
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
			expected: 3.1415,
		},
		{
			oldAngle: 0.0,
			newAngle: 3.15,
			expected: 3.14,
		},
		{
			oldAngle: 0.0,
			newAngle: 4.0,
			expected: 2.28,
		},
	}

	for _, tc := range testData {
		value := Difference(tc.oldAngle, tc.newAngle)

		diff := math.Abs(value - tc.expected)

		assert.Less(t, diff, float32(0.01))
	}
}

func TestDifferenceWithDirection(t *testing.T) {
	testData := []struct {
		oldAngle  float64
		newAngle  float64
		direction bool
		expected  float64
	}{
		{
			oldAngle:  0.0,
			newAngle:  0.0,
			direction: true,
			expected:  0.0,
		},
		{
			oldAngle:  0.0,
			direction: true,
			newAngle:  1.5,
			expected:  -1.5,
		},
		{
			oldAngle:  0.0,
			direction: true,
			newAngle:  4,
			expected:  2.28,
		},
		{
			oldAngle:  0.0,
			direction: false,
			newAngle:  0.0,
			expected:  0.0,
		},
		{
			oldAngle:  0.0,
			direction: false,
			newAngle:  1.5,
			expected:  1.5,
		},
		{
			oldAngle:  0.0,
			direction: false,
			newAngle:  4.0,
			expected:  -2.28,
		},
	}

	for _, tc := range testData {
		value := DifferenceWithDirection(tc.oldAngle, tc.newAngle, tc.direction)

		diff := math.Abs(value - tc.expected)

		assert.Less(t, diff, 0.01)
	}
}
