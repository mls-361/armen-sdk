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
		Devel() int
	}

	// Config AFAIRE.
	Config interface {
	}

	// Crypto AFAIRE.
	Crypto interface {
		EncryptString(s string) (string, error)
		DecryptString(s string) (string, error)
	}

	// Components AFAIRE.
	Components struct {
		Application Application
		Config      Config
		Crypto      Crypto
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
