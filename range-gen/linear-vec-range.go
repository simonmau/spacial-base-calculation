package rangegen

import "github.com/ungerik/go3d/float64/vec2"

func GenLinearVecRangeChan(minRange, maxRange *vec2.T, steps *float64, pipe chan<- vec2.T) {
	pipe <- *minRange

	diff := vec2.Sub(maxRange, minRange)
	diff.Scale(1.0 / *steps)

	pos := vec2.Add(minRange, &diff)

	for i := float64(0.0); i < *steps; i++ {
		pipe <- pos
		pos = vec2.Add(&pos, &diff)
	}

	pipe <- *maxRange

	close(pipe)
}
