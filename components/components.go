/*
------------------------------------------------------------------------------------------------------------------------
####### components ####### (c) 2020-2021 mls-361 ################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package components

import (
	"net/http"
	"time"

	"github.com/mls-361/datamap"

	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/armen-sdk/logger"
	"github.com/mls-361/armen-sdk/message"
)

type (
	// Component AFAIRE.
	Component interface {
		Name() string
	}

	// Application AFAIRE.
	Application interface {
		ID() string
		Name() string
		Version() string
		BuiltAt() time.Time
		StartedAt() time.Time
		LookupEnv(suffix string) (string, bool)
		Debug() int
		Host() string
	}

	// Bus AFAIRE.
	Bus interface {
		AddPublisher(name string, chCapacity, nbConsumer int) chan<- *message.Message
		Subscribe(callback func(*message.Message), regexpList ...string) error
	}

	// Config AFAIRE.
	Config interface {
		Data() datamap.DataMap
		Decode(to interface{}, mustExist bool, keys ...string) error
	}

	// Crypto AFAIRE.
	Crypto interface {
		EncryptString(s string) (string, error)
		DecryptString(s string) (string, error)
	}

	// Leader AFAIRE.
	Leader interface {
		AmITheLeader() bool
	}

	// Model AFAIRE.
	Model interface {
		AcquireLock(name, owner string, duration time.Duration) (bool, error)
		ReleaseLock(name, owner string) error
		InsertJob(job *jw.Job) error
		UpdateJob(job *jw.Job)
		InsertWorkflow(wf *jw.Workflow) error
	}

	// Router AFAIRE.
	Router interface {
		Get(path string, handler http.HandlerFunc)
		Post(path string, handler http.HandlerFunc)
	}

	// Components AFAIRE.
	Components interface {
		Application() Application
		Bus() Bus
		Config() Config
		Crypto() Crypto
		Leader() Leader
		Logger() logger.Logger
		Model() Model
		Router() Router
	}

	// Plugin AFAIRE.
	Plugin interface {
		ID() string
		Name() string
		Version() string
		BuiltAt() time.Time
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
