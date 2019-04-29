package toolscan

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

const expectedCmakeResult = "cmake version 3.14.3\n\nCMake suite maintained and supported by Kitware (kitware.com/cmake)"
const actualCmakeResult = "cmake version 3.14.3\n\nCMake suite maintained and supported by Kitware (kitware.com/cmake)"

//=============================================================================
func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	// some code here to check arguments perhaps?
	fmt.Fprintf(os.Stdout, actualCmakeResult)
	os.Exit(0)
}

func TestRunCmake(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	out, err := runner.RunCmake("--version")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedCmakeResult, string(out))
}

//=============================================================================
type runnerDummy struct {
	output []byte
	err    error
}

func (r runnerDummy) RunCmake(args ...string) ([]byte, error) {
	return r.output, r.err
}

func TestScanForCmake(t *testing.T) {
	runner = runnerDummy{
		output: []byte(actualCmakeResult),
		err:    nil,
	}
	defer func() { runner = newCommandlineRunner() }()
	toolInfo := scanForCmake()
	assert.Equal(t, "3.14.3", toolInfo.Version)
	assert.True(t, toolInfo.Available)
}

func TestScanForCmakeExecutionError(t *testing.T) {
	expectedError := errors.New("Execution error")
	runner = runnerDummy{
		output: nil,
		err:    expectedError,
	}
	defer func() { runner = newCommandlineRunner() }()
	_, err := runner.RunCmake("--version")
	assert.Error(t, expectedError, err)
}
