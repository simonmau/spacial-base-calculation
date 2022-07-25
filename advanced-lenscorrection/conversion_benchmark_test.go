package advancedlenscorrection

import (
	"testing"

	"github.com/ungerik/go3d/float64/vec2"
)

type testStruct struct {
	P      float64
	Point  vec2.T
	Center vec2.T
}

func BenchmarkConversions(b *testing.B) {
	b.StopTimer()

	item := testStruct{
		P:      0.31,
		Point:  vec2.T{302, 483},
		Center: vec2.T{860, 540},
	}

	width := float64(1920)
	height := float64(1080)

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sexyPt := CorrectRawToSexy(&item.P, &item.Point, &item.Center, &width, &height)

		if sexyPt != nil {
			_ = CorrectSexyToRaw(&item.P, sexyPt, &item.Center, &width, &height)
		}
	}
}
