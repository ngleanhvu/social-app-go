package main

import (
	"context"
	"crud-go/component/asyncjob"
	"log"
	"time"
)

func main() {

	//funcExecute := func(ctx context.Context) error {
	//	time.Sleep(5 * time.Second)
	//	log.Println("hello world")
	//	return errors.New("hello world error")
	//}
	//
	//job := asyncjob.NewJob(funcExecute)
	//
	//if err := job.Execute(context.Background()); err != nil {
	//	log.Println(job.State(), err)
	//	for {
	//		if err := job.Retry(context.Background()); err != nil {
	//			log.Println(job.State(), err)
	//		}
	//		if job.State() == asyncjob.StateRetryFailed || job.State() == asyncjob.StateCompleted {
	//			log.Println(job.State(), err)
	//			break
	//		}
	//	}
	//}

	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("I am job1")
		return nil
	})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 2)
		log.Println("I am job2")
		return nil
	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 3)
		log.Println("I am job3")
		return nil
	})

	group := asyncjob.NewGroup(false, job1, job2, job3)

	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}

}
