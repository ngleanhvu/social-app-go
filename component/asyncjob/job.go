package asyncjob

import (
	"context"
	"time"
)

type JobState int

type Job interface {
	Execute(ctx context.Context) error
	Retry(context.Context) error
	JobState() JobState
	SetRetryDurations(times []time.Duration)
}

const (
	defaultMaxTimeout = time.Second * 10
)

var (
	defaultRetryTime = []time.Duration{
		time.Second,
		time.Second * 5,
		time.Second * 10,
	}
)

type JobHandler func(ctx context.Context) error

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(jobHandler JobHandler) *job {
	j := job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    jobHandler,
		state:      StateInit,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}
	return &j
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning

	var err error
	err = j.handler(ctx)

	if err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted
	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])

	err := j.handler(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed
	return err
}

func (j *job) JobState() JobState {
	return j.state
}

func (j *job) RetryIndex() int {
	return j.retryIndex
}

func (j *job) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}
	j.config.Retries = times
}
