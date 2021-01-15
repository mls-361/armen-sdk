/*
------------------------------------------------------------------------------------------------------------------------
####### logger ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package logger

import (
	"log"

	"github.com/mls-361/logger"
)

type (
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
)

/*
######################################################################################################## @(°_°)@ #######
*/
