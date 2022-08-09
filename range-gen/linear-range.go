package rangegen

func GenLinearRangeChan(minRange, maxRange, steps *float64, pipe chan<- float64) {
	pipe <- *minRange

	diff := *maxRange - *minRange
	stepSize := diff / *steps

	for i := *minRange + stepSize; i < *maxRange; i += stepSize {
		pipe <- i
	}

	pipe <- *maxRange

	close(pipe)
}

func GenLinearRangeArray(minRange, maxRange, steps *float64) []float64 {
	resArray := make([]float64, 0)

	resArray = append(resArray, *minRange)

	diff := *maxRange - *minRange
	stepSize := diff / *steps

	max := *maxRange - (stepSize * 0.5)

	for i := *minRange + stepSize; i < max; i += stepSize {
		resArray = append(resArray, i)
	}

	resArray = append(resArray, *maxRange)

	return resArray
}
