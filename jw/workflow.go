/*
------------------------------------------------------------------------------------------------------------------------
####### jw ####### (c) 2020-2021 mls-361 ########################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package jw

import (
	"time"
)

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
