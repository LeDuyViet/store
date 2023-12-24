package main

import (
	"sync"
)

type WorkerPool struct {
	NumWorkers int
	JobCh      chan interface{}
	ResultCh   chan interface{}
	ProcessJob func(job interface{}, resultCh chan interface{}, JobCh chan interface{})
}

func (wp *WorkerPool) Run() {
	var wg sync.WaitGroup
	for i := 0; i < wp.NumWorkers; i++ {
		wg.Add(1)
		worker := Worker{ID: i + 1, JobCh: wp.JobCh, Result: make(chan interface{})}
		go func() {
			defer wg.Done()
			for job := range worker.JobCh {
				wp.ProcessJob(job, wp.ResultCh, wp.JobCh)
			}
		}()
	}
	wg.Wait()
	// close(wp.ResultCh)
}
