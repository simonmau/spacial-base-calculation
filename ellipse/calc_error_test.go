package ellipse

import (
	"fmt"
	"testing"

	"github.com/ungerik/go3d/float64/vec2"

	"github.com/loov/hrtime"
	"github.com/loov/hrtime/hrtesting"
	"github.com/stretchr/testify/assert"
)

func TestCalcError(t *testing.T) {
	t.Parallel()

	data := Data{
		A: 1.0,
		B: 0.0,
		C: 1.0,
		D: 0.0,
		E: 0.0,
		F: -100.0,
	}

	errorList := []vec2.T{
		{5.0, 15.0},
		{5.0, 10.0},
		{5.0, 5.0},
		{5.0, 0.0},
		{5.0, -5.0},
		{5.0, -10.0},
		{5.0, -15.0},
	}

	res := data.CalcAverageError(&errorList)

	assert.NotNil(t, res)
}

func BenchmarkCalcError(b *testing.B) {
	data := Data{
		A: 1.0,
		B: 0.0,
		C: 1.0,
		D: 0.0,
		E: 0.0,
		F: -100.0,
	}

	errorList := []vec2.T{
		{5.0, 15.0},
		{5.0, 10.0},
		{5.0, 5.0},
		{5.0, 0.0},
		{5.0, -5.0},
		{5.0, -10.0},
		{5.0, -15.0},
	}

	for i := 0; i < b.N; i++ {
		for y := 0; y < 10000; y++ {
			_ = data.CalcAverageError(&errorList)
		}
	}
}

func BenchmarkCalcErrorHistogram(b *testing.B) {
	bench := hrtesting.NewBenchmark(b)
	benchHist := hrtime.NewBenchmark(b.N)
	defer bench.Report()

	data := Data{
		A: 1.0,
		B: 0.0,
		C: 1.0,
		D: 0.0,
		E: 0.0,
		F: -100.0,
	}

	errorList := []vec2.T{
		{5.0, 15.0},
		{5.0, 10.0},
		{5.0, 5.0},
		{5.0, 0.0},
		{5.0, -5.0},
		{5.0, -10.0},
		{5.0, -15.0},
	}

	for bench.Next() {
		for i := 0; i < 10000; i++ {
			_ = data.CalcAverageError(&errorList)
		}
	}

	for benchHist.Next() {
		for i := 0; i < 10000; i++ {
			_ = data.CalcAverageError(&errorList)
		}
	}

	fmt.Println(benchHist.Histogram(10))
}

func TestCalculateSchnitt(t *testing.T) {
	data := Data{
		A: 1.0,
		B: 0.0,
		C: 1.0,
		D: 0.0,
		E: 0.0,
		F: -100.0,
	}

	array := data.CalculateY(0)

	assert.NotNil(t, array)
	assert.NotEmpty(t, array)
	assert.InDelta(t, 0, array[0][0], 0.0001)
	assert.InDelta(t, -10, array[0][1], 0.0001)
	assert.InDelta(t, 0, array[1][0], 0.0001)
	assert.InDelta(t, 10, array[1][1], 0.0001)
}

func TestCalculateSchnitt2(t *testing.T) {
	data := Data{
		A: 0.69950384,
		B: 0.000836124,
		C: 1.0,
		D: -1326.124,
		E: -566.7868,
		F: 681281.1,
	}

	array := data.CalculateY(920.0)

	assert.NotNil(t, array)
	assert.NotEmpty(t, array)
	assert.InDelta(t, 920, array[0][0], 0.0001)
	assert.InDelta(t, 119.341705, array[0][1], 0.0001)
	assert.InDelta(t, 920, array[1][0], 0.0001)
	assert.InDelta(t, 446.67587, array[1][1], 0.0001)
}
