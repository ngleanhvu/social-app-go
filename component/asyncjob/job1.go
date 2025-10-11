package asyncjob

import (
	"context"
	"time"
)

type Job1 interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	JobState1() JobState1
	SetRetryDurations(times []time.Duration)
}

const (
	defaultMaxTimeout1 = time.Second * 10
)

var (
	defaultRetryTime1 = []time.Duration{
		time.Second,
		time.Second * 5,
		time.Second * 10,
	}
)

type JobState1 int

const (
	StateInit1 JobState1 = iota
	StateRunning1
	StateCompleted1
	StateFailed1
	StateRetryFailed1
)

func (js JobState1) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type jobConfig1 struct {
	MaxTimeout1 time.Duration
	Retries1    []time.Duration
}

type JobHandler1 func(ctx context.Context) error

type job1 struct {
	jobConfig  jobConfig1
	jobHandler JobHandler1
	jobState   JobState1
	retryIndex int
	stopCh     chan bool
}

func NewJob1(handler JobHandler1) *job1 {
	return &job1{
		jobConfig: jobConfig1{
			MaxTimeout1: defaultMaxTimeout1,
			Retries1:    defaultRetryTime1,
		},
		jobHandler: handler,
		jobState:   StateInit1,
		retryIndex: 0,
		stopCh:     make(chan bool),
	}
}

func (j *job1) Execute(ctx context.Context) error {
	j.jobState = StateRunning1

	var err error
	err = j.jobHandler(ctx)

	if err != nil {
		j.jobState = StateFailed1
		return err
	}

	j.jobState = StateCompleted1
	return nil
}

func (j *job1) Retry(ctx context.Context) error {
	j.retryIndex += 1
	time.Sleep(j.jobConfig.Retries1[j.retryIndex])

	err := j.jobHandler(ctx)

	if err == nil {
		j.jobState = StateCompleted1
		return nil
	}

	if j.retryIndex == len(j.jobConfig.Retries1)-1 {
		j.jobState = StateRetryFailed1
		return err
	}

	j.jobState = StateFailed1
	return err
}

func (j *job1) JobState1() JobState1 {
	return j.jobState
}

func (j *job1) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}
	j.jobConfig.Retries1 = times
}
