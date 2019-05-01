package toolscan

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const actualCmakeResult = "cmake version 3.14.3\n\nCMake suite maintained and supported by Kitware (kitware.com/cmake)\n"
const actualCppCheckResult = "Cppcheck 1.86\n"

//=============================================================================
func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", strings.Join([]string{"COMMAND=", command}, "")}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	var cmd = os.Getenv("COMMAND")
	var result string
	switch cmd {
	case "cmake":
		result = actualCmakeResult
	case "cppcheck":
		result = actualCppCheckResult
	}

	// some code here to check arguments perhaps?
	fmt.Fprintf(os.Stdout, result)
	os.Exit(0)
}

//=============================================================================
func TestRunCmake(t *testing.T) {
	const expectedCmakeResult = "cmake version 3.14.3\n\nCMake suite maintained and supported by Kitware (kitware.com/cmake)\n"
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	out, err := runner.RunCmake("--version")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedCmakeResult, string(out))
}

func TestRunCppCheck(t *testing.T) {
	const expectedCppCheckResult = "Cppcheck 1.86\n"
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	out, err := runner.RunCppCheck("--version")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedCppCheckResult, string(out))
}

//=============================================================================
type runnerDummy struct {
	output []byte
	err    error
}

func (r runnerDummy) RunCmake(args ...string) ([]byte, error) {
	return r.output, r.err
}

func (r runnerDummy) RunCppCheck(args ...string) ([]byte, error) {
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
	toolInfo := scanForCmake()
	assert.False(t, toolInfo.Available)
	assert.Equal(t, "", toolInfo.Version)
}

func TestScanForCppCheck(t *testing.T) {
	runner = runnerDummy{
		output: []byte(actualCppCheckResult),
		err:    nil,
	}
	defer func() { runner = newCommandlineRunner() }()
	toolInfo := scanForCppCheck()
	assert.True(t, toolInfo.Available)
	assert.Equal(t, "1.86", toolInfo.Version)
}

func TestScanForCppCheckExecutionError(t *testing.T) {
	expectedError := errors.New("Execution error")
	runner = runnerDummy{
		output: nil,
		err:    expectedError,
	}
	defer func() { runner = newCommandlineRunner() }()
	toolInfo := scanForCppCheck()
	assert.False(t, toolInfo.Available)
	assert.Equal(t, "", toolInfo.Version)
}
