/*
------------------------------------------------------------------------------------------------------------------------
####### jw ####### (c) 2020-2021 mls-361 ########################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package jw

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

/*
######################################################################################################## @(°_°)@ #######
*/
