package drawingprecision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStepsCalculation_1(t *testing.T) {
	dp := GenDrawPrecByDistance(0.1)
	steps := dp.CalcSteps(1)
	assert.InDelta(t, 10, steps, 0.0001)
}

func TestStepsCalculation_2(t *testing.T) {
	dp := GenDrawPrecByPointsPerUnit(10)
	steps := dp.CalcSteps(1)
	assert.InDelta(t, 10, steps, 0.0001)
}
