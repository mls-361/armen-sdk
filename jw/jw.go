/*
------------------------------------------------------------------------------------------------------------------------
####### jw ####### (c) 2020-2021 mls-361 ########################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package jw

import (
	"time"

	"github.com/mls-361/datamap"

	"github.com/mls-361/armen-sdk/logger"
)

type (
	// Priority AFAIRE.
	Priority int
	// Exclusivity AFAIRE.
	Exclusivity string
	// Status AFAIRE.
	Status string
)

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
		Config         datamap.DataMap
		Data           datamap.DataMap
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
func NewJob(id, name, namespace, _type, origin string, priority Priority, key, emails *string) *Job {
	now := time.Now()

	return &Job{
		ID:            id,
		Name:          name,
		Namespace:     namespace,
		Type:          _type,
		Origin:        origin,
		Priority:      priority,
		Key:           key,
		Emails:        emails,
		Config:        make(datamap.DataMap),
		Data:          make(datamap.DataMap),
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

/*
######################################################################################################## @(°_°)@ #######
*/
