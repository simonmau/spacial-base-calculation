package rangegen

import "github.com/ungerik/go3d/float64/vec2"

func GenSquareVecRangeChan(topLeft, bottomRight *vec2.T, steps *float64, pipe chan<- vec2.T) {
	//PERF: is gen-linear-range-array faster?

	xChan := make(chan float64)

	go GenLinearRangeChan(&topLeft[0], &bottomRight[0], steps, xChan)

	for x := range xChan {
		yChan := make(chan float64)

		go GenLinearRangeChan(&topLeft[1], &bottomRight[1], steps, yChan)

		for y := range yChan {
			pipe <- vec2.T{x, y}
		}
	}

	close(pipe)
}

func GenSquareVecRangeArray(topLeft, bottomRight *vec2.T, steps *float64) []vec2.T {
	resArray := make([]vec2.T, 0)

	xRange := GenLinearRangeArray(&topLeft[0], &bottomRight[0], steps)
	yRange := GenLinearRangeArray(&topLeft[1], &bottomRight[1], steps)

	for _, x := range xRange {
		for _, y := range yRange {
			resArray = append(resArray, vec2.T{x, y})
		}
	}

	return resArray
}
