package parallelwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	data := make([]int, 0)
	data = append(data, 0)
	data = append(data, 1)
	data = append(data, 2)
	data = append(data, 3)
	data = append(data, 4)
	data = append(data, 5)

	service := CreateService[int, int](&data, 1)

	service.Run(func(input *int) int {
		val := *input
		return val * val
	})

	resChannel := service.GetResultChannel()

	for i := 0; i < len(data); i++ {
		item := <-resChannel

		assert.Equal(t, i*i, item)
	}
}
