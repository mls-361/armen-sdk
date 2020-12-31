/*
------------------------------------------------------------------------------------------------------------------------
####### components ####### (c) 2020-2021 mls-361 ################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package components

import (
	"log"
	"time"

	"github.com/mls-361/datamap"
	"github.com/mls-361/logger"

	"github.com/mls-361/armen-sdk/message"
)

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

	// Components AFAIRE.
	Components struct {
		Application Application
		Bus         Bus
		Config      Config
		Crypto      Crypto
		Logger      Logger
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
