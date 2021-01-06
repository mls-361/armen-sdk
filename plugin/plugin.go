/*
------------------------------------------------------------------------------------------------------------------------
####### plugin ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package plugin

import (
	"strconv"
	"time"

	"github.com/mls-361/minikit"
	"github.com/mls-361/uuid"

	"github.com/mls-361/armen-sdk/components"
)

type (
	// Plugin AFAIRE.
	Plugin struct {
		*minikit.Base
		id         string
		name       string
		version    string
		builtAt    time.Time
		components *components.Components
	}
)

// New AFAIRE.
func New(name, version, builtAt string, components *components.Components) *Plugin {
	ts, err := strconv.ParseInt(builtAt, 0, 64)
	if err != nil {
		ts = 0
	}

	plugin := &Plugin{
		Base:       minikit.NewBase(name, "plugin."+name),
		id:         uuid.New(),
		name:       name,
		version:    version,
		builtAt:    time.Unix(ts, 0),
		components: components,
	}

	// Pour les besoins de l'application.
	components.Add(plugin)

	return plugin
}

// Dependencies AFAIRE.
func (cp *Plugin) Dependencies() []string {
	return []string{
		"logger",
	}
}

// Build AFAIRE.
func (cp *Plugin) Build(_ *minikit.Manager) error {
	cp.components.Logger.Info( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"Plugin",
		"name", cp.name,
		"version", cp.version,
		"builtAt", cp.builtAt.String(),
	)

	return nil
}

// ID AFAIRE.
func (cp *Plugin) ID() string {
	return cp.id
}

// Name AFAIRE.
func (cp *Plugin) Name() string {
	return cp.name
}

// Version AFAIRE.
func (cp *Plugin) Version() string {
	return cp.version
}

// BuiltAt AFAIRE.
func (cp *Plugin) BuiltAt() time.Time {
	return cp.builtAt
}

/*
######################################################################################################## @(°_°)@ #######
*/
