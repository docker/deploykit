package launch

import (
	log "github.com/Sirupsen/logrus"
)

type noOp int

// NewNoOpLauncher doesn't actually launch the plugins.  It's a stub with no op and relies on manual plugin starts.
func NewNoOpLauncher() Launcher {
	return noOp(0)
}

// Name returns the name of the launcher
func (n noOp) Name() string {
	return "noop"
}

// Launch starts the plugin given the name
func (n noOp) Launch(name string, config *Config) (<-chan error, error) {
	log.Infoln("NO-OP Launcher: not automatically starting plugin", name, "args=", config)

	starting := make(chan error)
	close(starting) // channel won't block
	return starting, nil
}
