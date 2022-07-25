package parallelwork

import (
	"errors"
	"runtime"
)

type ParallelWorkService[I any, O any] struct {
	inputChannel          chan I
	workerFinishedChannel chan struct{}
	resultChannel         chan O

	numberOfThreads int
	finished        bool
}

//creates a new parallel-work-service for the input, if number-of-threads is < 1 it will be automatically set
func CreateService[I any, O any](input *[]I, numberOfThreads int) ParallelWorkService[I, O] {
	if numberOfThreads < 1 {
		numberOfThreads = runtime.GOMAXPROCS(0) * 2
	}

	inputChan := make(chan I, numberOfThreads*3)
	workerFinishedChan := make(chan struct{})
	resultChan := make(chan O, len(*input))

	//push calc-data async
	go func() {
		for _, item := range *input {
			inputChan <- item
		}

		close(inputChan)
	}()

	return ParallelWorkService[I, O]{
		inputChannel:          inputChan,
		workerFinishedChannel: workerFinishedChan,
		resultChannel:         resultChan,
		numberOfThreads:       numberOfThreads,
		finished:              false,
	}
}

func (s *ParallelWorkService[I, O]) Run(function func(input *I) O) error {
	if s.finished {
		return errors.New("work-set is already finished, cannot run again")
	}

	for i := 0; i < s.numberOfThreads; i++ {
		//start workers asynchron
		go s.runParallelWorkerThread(function)
	}

	go s.closeDataChannelOnWorkerFinish()

	return nil
}

func (s *ParallelWorkService[I, O]) GetResultChannel() <-chan O {
	return s.resultChannel
}

func (s *ParallelWorkService[I, O]) runParallelWorkerThread(function func(input *I) O) {
	for item := range s.inputChannel {
		current := item
		result := function(&current)
		s.resultChannel <- result
	}

	s.workerFinishedChannel <- struct{}{}
}

func (s *ParallelWorkService[I, O]) closeDataChannelOnWorkerFinish() {
	for i := 0; i < s.numberOfThreads; i++ {
		<-s.workerFinishedChannel
	}

	close(s.resultChannel)
	close(s.workerFinishedChannel)

	s.finished = true
}
