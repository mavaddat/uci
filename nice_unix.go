//go:build unix

package uci

import "syscall"

// NewEngineNice creates a new Engine with the specified nice level.
// The nice value ranges from -20 (highest priority) to 19 (lowest priority).
// Note: Setting negative nice values typically requires root privileges.
func NewEngineNice(nice int, path string, arg ...string) (*Engine, error) {
	eng, err := NewEngine(path, arg...)
	if err != nil {
		return nil, err
	}
	err = syscall.Setpriority(syscall.PRIO_PROCESS, eng.cmd.Process.Pid, nice)
	if err != nil {
		eng.Close()
		return nil, err
	}
	return eng, nil
}

// SetNice changes the nice level of a running engine process.
// The nice value ranges from -20 (highest priority) to 19 (lowest priority).
// Note: Setting negative nice values typically requires root privileges.
func (eng *Engine) SetNice(nice int) error {
	return syscall.Setpriority(syscall.PRIO_PROCESS, eng.cmd.Process.Pid, nice)
}
