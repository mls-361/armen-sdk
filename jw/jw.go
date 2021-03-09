/*
------------------------------------------------------------------------------------------------------------------------
####### jw ####### (c) 2020-2021 mls-361 ########################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package jw

import (
	"time"

	"github.com/mls-361/logger"
)

type (
	// Priority AFAIRE.
	Priority int
	// Status AFAIRE.
	Status string
)

// String AFAIRE.
func (s Status) String() string { return string(s) }

const (
	// NoPriority AFAIRE.
	NoPriority Priority = 0
	// PriorityLow AFAIRE.
	PriorityLow Priority = 20
	// PriorityMedium AFAIRE.
	PriorityMedium Priority = 50
	// PriorityHigh AFAIRE.
	PriorityHigh Priority = 80

	// StatusToDo AFAIRE.
	StatusToDo Status = "todo"
	// StatusRunning AFAIRE.
	StatusRunning Status = "running"
	// StatusPending AFAIRE.
	StatusPending Status = "pending"
	// StatusSucceeded AFAIRE.
	StatusSucceeded Status = "succeeded"
	// StatusFailed AFAIRE.
	StatusFailed Status = "failed"
)

type (
	// Job AFAIRE.
	Job struct {
		ID             string
		Name           string
		Namespace      string
		Type           string
		Origin         string
		Priority       Priority
		Key            *string
		Workflow       *string
		WorkflowFailed *bool
		Emails         *string
		Config         map[string]interface{}
		Private        map[string]interface{}
		Public         map[string]interface{}
		CreatedAt      time.Time
		Status         Status
		Error          *string
		Attempts       int
		FinishedAt     *time.Time
		RunAfter       time.Time
		Result         *string
		NextStep       *string
		Weight         int
		TimeReference  time.Time
	}

	// Result AFAIRE.
	Result struct {
		Status   Status
		Err      error
		Duration time.Duration
	}

	// Runner AFAIRE.
	Runner interface {
		RunJob(job *Job, logger logger.Logger) *Result
	}
)

// NewJob AFAIRE.
func NewJob(id, name, namespace, _type, origin string, priority Priority, key *string) *Job {
	now := time.Now()

	return &Job{
		ID:            id,
		Name:          name,
		Namespace:     namespace,
		Type:          _type,
		Origin:        origin,
		Priority:      priority,
		Key:           key,
		CreatedAt:     now,
		Status:        StatusToDo,
		RunAfter:      now,
		TimeReference: now,
	}
}

func newResult(status Status, err error, duration time.Duration) *Result {
	return &Result{
		Status:   status,
		Err:      err,
		Duration: duration,
	}
}

// Succeeded AFAIRE.
func Succeeded(err error) *Result {
	return newResult(StatusSucceeded, err, 0)
}

// Failed AFAIRE.
func Failed(err error) *Result {
	return newResult(StatusFailed, err, 0)
}

// Pending AFAIRE.
func Pending(err error, duration time.Duration) *Result {
	return newResult(StatusPending, err, duration)
}

type (
	// Step AFAIRE.
	Step struct {
		Namespace string
		Type      string
		Config    map[string]interface{}
		Next      map[string]interface{}
	}

	// Workflow AFAIRE.
	Workflow struct {
		ID                string
		Type              string
		Description       string
		Origin            string
		Priority          Priority
		FirstStep         string
		AllSteps          map[string]*Step
		ExternalReference *string
		Emails            *string
		Data              map[string]interface{}
		CreatedAt         time.Time
		Status            Status
		FinishedAt        *time.Time
	}
)

// NewStep AFAIRE.
func NewStep(namespace, _type string, config, next map[string]interface{}) *Step {
	return &Step{
		Namespace: namespace,
		Type:      _type,
		Config:    config,
		Next:      next,
	}
}

// NewWorkflow AFAIRE.
func NewWorkflow(id, _type, desc, origin string, priority Priority, fStep string, aSteps map[string]*Step) *Workflow {
	return &Workflow{
		ID:          id,
		Type:        _type,
		Description: desc,
		Origin:      origin,
		Priority:    priority,
		FirstStep:   fStep,
		AllSteps:    aSteps,
		CreatedAt:   time.Now(),
		Status:      StatusRunning,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
