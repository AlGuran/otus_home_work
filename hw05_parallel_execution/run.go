package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	taskChan := make(chan Task, len(tasks))
	errChan := make(chan error, m+n)

	for _, t := range tasks {
		taskChan <- t
	}
	close(taskChan)

	wg := new(sync.WaitGroup)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(taskChan, errChan, wg, m)
	}
	wg.Wait()
	close(errChan)

	if len(errChan) >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func worker(taskCh <-chan Task, errChan chan<- error, wg *sync.WaitGroup, errCnt int) {
	defer wg.Done()

	for {
		if len(errChan) >= errCnt {
			return
		}

		task, ok := <-taskCh
		if !ok {
			return
		}

		err := task()
		if err != nil {
			errChan <- err
		}
	}
}
