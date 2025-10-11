package asyncjob

import (
	"context"
	"log"
	"sync"
)

type group1 struct {
	job1s        []Job1
	isConcurrent bool
	wg           *sync.WaitGroup
}

func newGroup1(isConcurrent bool, job1s ...Job1) *group1 {
	return &group1{
		job1s:        job1s,
		isConcurrent: isConcurrent,
		wg:           new(sync.WaitGroup),
	}
}

func (g *group1) Run1(ctx context.Context) error {
	g.wg.Add(len(g.job1s))
	errChan := make(chan error, len(g.job1s))

	for i, _ := range g.job1s {
		job := g.job1s[i]
		if err := g.runJob1(ctx, job); err != nil {
			errChan <- err
		}
	}

	var err error

	for i := 0; i < len(g.job1s); i++ {
		if v := <-errChan; v != nil {
			err = v
		}
	}

	return err
}

func (g *group1) runJob1(ctx context.Context, job1 Job1) error {
	if err := job1.Execute(ctx); err != nil {
		log.Println(err)
		for {
			if job1.JobState1() == StateRetryFailed1 {
				log.Println(err)
				return err
			}

			if job1.Retry(ctx) == nil {
				return nil
			}
		}
	}
	return nil
}
