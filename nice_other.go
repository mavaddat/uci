//go:build !unix

package uci

import "errors"

var errNiceNotSupported = errors.New("nice level is not supported on this platform")

// NewEngineNice creates a new Engine with the specified nice level.
// This function is not supported on non-Unix platforms and will return an error.
func NewEngineNice(nice int, path string, arg ...string) (*Engine, error) {
	return nil, errNiceNotSupported
}

// SetNice changes the nice level of a running engine process.
// This function is not supported on non-Unix platforms and will return an error.
func (eng *Engine) SetNice(nice int) error {
	return errNiceNotSupported
}
