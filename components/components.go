/*
------------------------------------------------------------------------------------------------------------------------
####### components ####### (c) 2020-2021 mls-361 ################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package components

import "time"

type (
	// Application AFAIRE.
	Application interface {
		ID() string
		Name() string
		Version() string
		BuiltAt() time.Time
		StartedAt() time.Time
		Host() string
		Devel() bool
	}

	// Components AFAIRE.
	Components struct {
		Application Application
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
