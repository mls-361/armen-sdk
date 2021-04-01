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

	// JobCore AFAIRE.
	JobCore struct {
		Name      string                 `json:"Name"`
		Namespace string                 `json:"Namespace"`
		Type      string                 `json:"type"`
		Priority  string                 `json:"priority"`
		Key       string                 `json:"key"`
		Data      map[string]interface{} `json:"data"`
	}

	// Factory AFAIRE.
	Factory interface {
		CreateJob(jc *JobCore) (*Job, error)
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

/*
######################################################################################################## @(°_°)@ #######
*/
