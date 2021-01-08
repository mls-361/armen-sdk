/*
------------------------------------------------------------------------------------------------------------------------
####### worker ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package worker

import (
	"time"

	"github.com/mls-361/uuid"
)

type (
	// Worker AFAIRE.
	Worker struct {
		ID        string
		StartedAt time.Time
		Jobs      int
		Data      interface{}
	}
)

// New AFAIRE.
func New() *Worker {
	return &Worker{
		ID:        uuid.New(),
		StartedAt: time.Now(),
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/