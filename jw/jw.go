/*
------------------------------------------------------------------------------------------------------------------------
####### jw ####### (c) 2020-2021 mls-361 ########################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package jw

import (
	"time"

	"github.com/mls-361/datamap"
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

	// NoExclusivity AFAIRE.
	NoExclusivity Exclusivity = "no"
	// ExclusiveWithItself AFAIRE.
	ExclusiveWithItself Exclusivity = "itself"
	// ApplicationLevelExclusivity AFAIRE.
	ApplicationLevelExclusivity Exclusivity = "application"

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
		Exclusivity    Exclusivity
		UniqueKey      *string
		Emails         *string
		Config         datamap.DataMap
		Data           datamap.DataMap
		Workflow       *string
		WorkflowFailed *bool
		CreatedAt      time.Time
		TimeReference  time.Time
		Status         Status
		Error          *string
		Attempts       int
		RunAfter       time.Time
		Result         *string
		NextStep       *string
		FinishedAt     *time.Time
	}

	// Result AFAIRE.
	Result struct {
		Status   Status
		Err      error
		Duration time.Duration
	}
)

// NewJob AFAIRE.
func NewJob(id, n, ns, t, o string, p Priority, e Exclusivity, uk, em *string) *Job {
	now := time.Now()

	return &Job{
		ID:            id,
		Name:          n,
		Namespace:     ns,
		Type:          t,
		Origin:        o,
		Priority:      p,
		Exclusivity:   e,
		UniqueKey:     uk,
		Emails:        em,
		Config:        make(datamap.DataMap),
		Data:          make(datamap.DataMap),
		CreatedAt:     now,
		TimeReference: now,
		Status:        StatusToDo,
		RunAfter:      now,
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
