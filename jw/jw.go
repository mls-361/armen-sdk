/*
------------------------------------------------------------------------------------------------------------------------
####### jw ####### (c) 2020-2021 mls-361 ########################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package jw

import (
	"strconv"

	"github.com/mls-361/failure"
)

type (
	// Priority AFAIRE.
	Priority int
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
	// PriorityCritical AFAIRE.
	PriorityCritical Priority = 100

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

// IntToPriority AFAIRE.
func IntToPriority(priority int) Priority {
	if priority > int(PriorityCritical) {
		return PriorityCritical
	}

	if priority < int(NoPriority) {
		return NoPriority
	}

	return Priority(priority)
}

// StringToPriority AFAIRE.
func StringToPriority(priority string) (Priority, error) {
	switch priority {
	case "none", "None", "NONE":
		return NoPriority, nil
	case "low", "Low", "LOW":
		return PriorityLow, nil
	case "medium", "Medium", "MEDIUM":
		return PriorityMedium, nil
	case "high", "High", "HIGH":
		return PriorityHigh, nil
	case "critical", "Critical", "CRITICAL":
		return PriorityCritical, nil
	}

	p, err := strconv.Atoi(priority)
	if err != nil {
		return 0, failure.New(err).
			Set("value", priority).
			Msg("this value is not a priority") ////////////////////////////////////////////////////////////////////////
	}

	return IntToPriority(p), nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
