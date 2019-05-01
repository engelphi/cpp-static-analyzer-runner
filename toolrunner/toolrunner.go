package toolrunner

import "os/exec"

var execCommand = exec.Command

// CommandlineRunner exectues commandline tools and returns the output
type CommandlineRunner interface {
	RunCmake(arg ...string) (out []byte, err error)
	RunCppCheck(arg ...string) (out []byte, err error)
}

type commandlineRunnerImpl struct{}

// RunCmake runs cmake with the given arguments and returns the output
func (r commandlineRunnerImpl) RunCmake(arg ...string) (out []byte, err error) {
	cmd := execCommand("cmake", arg...)
	out, err = cmd.CombinedOutput()
	return
}

func (r commandlineRunnerImpl) RunCppCheck(arg ...string) (out []byte, err error) {
	cmd := execCommand("cppcheck", arg...)
	out, err = cmd.CombinedOutput()
	return
}

// NewCommandlineRunner consructs a new CommandlineRunner
func NewCommandlineRunner() CommandlineRunner {
	return commandlineRunnerImpl{}
}
