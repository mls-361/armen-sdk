/*
------------------------------------------------------------------------------------------------------------------------
####### components ####### (c) 2020-2021 mls-361 ################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package components

import (
	"log"
	"net/http"
	"time"

	"github.com/mls-361/datamap"
	"github.com/mls-361/logger"

	"github.com/mls-361/armen-sdk/jw"
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

	// Backend AFAIRE.
	Backend interface {
		AcquireLock(name, owner string, duration time.Duration) (bool, error)
		ReleaseLock(name, owner string) error
		InsertJob(job *jw.Job) (bool, error)
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
		Start()
		Stop()
	}

	// Logger AFAIRE.
	Logger interface {
		Trace(msg string, data ...interface{})
		Debug(msg string, data ...interface{})
		Info(msg string, data ...interface{})
		Notice(msg string, data ...interface{})
		Warning(msg string, data ...interface{})
		Error(msg string, data ...interface{})
		Fatal(msg string, data ...interface{})
		SetLevel(level string)
		CreateLogger(id, name string) *logger.Logger // BOF?
		RemoveLogger(id string)
		NewStdLogger(level, prefix string, flag int) *log.Logger
	}

	// Metrics AFAIRE.
	Metrics interface {
		AddInt(value int64, keys ...string)
		SetInt(value int64, keys ...string)
	}

	// Model AFAIRE.
	Model interface {
		AcquireLock(name, owner string, duration time.Duration) (bool, error)
		ReleaseLock(name, owner string) error
		InsertJob(job *jw.Job) error
		NextJob() *jw.Job
	}

	// Router AFAIRE.
	Router interface {
		Handler() http.Handler
		Get(path string, handler http.HandlerFunc)
		Post(path string, handler http.HandlerFunc)
	}

	// Scheduler AFAIRE.
	Scheduler interface {
		Start()
		Stop()
	}

	// Server AFAIRE.
	Server interface {
		Port() int
		Start() error
		Stop()
	}

	// Workers AFAIRE.
	Workers interface {
		Start()
		Stop()
	}

	// Components AFAIRE.
	Components struct {
		Application Application
		Backend     Backend
		Bus         Bus
		Config      Config
		Crypto      Crypto
		Leader      Leader
		Logger      Logger
		Metrics     Metrics
		Model       Model
		Router      Router
		Scheduler   Scheduler
		Server      Server
		Workers     Workers
		unknown     []Component
	}

	// Plugin AFAIRE.
	Plugin interface {
		ID() string
		Name() string
		Version() string
		BuiltAt() time.Time
	}
)

// New AFAIRE.
func New() *Components {
	return &Components{
		unknown: make([]Component, 0),
	}
}

// Add AFAIRE.
func (cs *Components) Add(c Component) {
	cs.unknown = append(cs.unknown, c)
}

/*
######################################################################################################## @(°_°)@ #######
*/
