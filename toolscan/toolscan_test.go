package toolscan

import (
	"errors"
	"strings"
	"testing"

	"github.com/engelphi/cpp-static-analyzer-runner/toolrunner"
	"github.com/stretchr/testify/assert"
)

const actualCmakeResult = "cmake version 3.14.3\n\nCMake suite maintained and supported by Kitware (kitware.com/cmake)\n"
const actualCppCheckResult = "Cppcheck 1.86\n"

//=============================================================================
type runnerDummy struct {
	outputCmake    []byte
	outputCppCheck []byte
	errCmake       error
	errCppCheck    error
}

func (r runnerDummy) RunCmake(args ...string) ([]byte, error) {
	return r.outputCmake, r.errCmake
}

func (r runnerDummy) RunCppCheck(args ...string) ([]byte, error) {
	return r.outputCppCheck, r.errCppCheck
}

//=============================================================================
func TestScanForCmake(t *testing.T) {
	runner = runnerDummy{
		outputCmake: []byte(actualCmakeResult),
		errCmake:    nil,
	}
	defer func() { runner = toolrunner.NewCommandlineRunner() }()
	toolInfo := scanForCmake()
	assert.Equal(t, "3.14.3", toolInfo.Version)
	assert.True(t, toolInfo.Available)
}

func TestScanForCmakeExecutionError(t *testing.T) {
	expectedError := errors.New("Execution error")
	runner = runnerDummy{
		outputCmake: nil,
		errCmake:    expectedError,
	}
	defer func() { runner = toolrunner.NewCommandlineRunner() }()
	toolInfo := scanForCmake()
	assert.False(t, toolInfo.Available)
	assert.Equal(t, "", toolInfo.Version)
}

func TestScanForCmakeInvalidToolOutput(t *testing.T) {
	runner = runnerDummy{
		outputCmake: []byte("cmake version 3.14.3"),
		errCmake:    nil,
	}
	defer func() { runner = toolrunner.NewCommandlineRunner() }()
	toolInfo := scanForCmake()
	assert.False(t, toolInfo.Available)
	assert.Equal(t, "", toolInfo.Version)
}

//=============================================================================
func TestScanForCppCheck(t *testing.T) {
	runner = runnerDummy{
		outputCppCheck: []byte(actualCppCheckResult),
		errCppCheck:    nil,
	}
	defer func() { runner = toolrunner.NewCommandlineRunner() }()
	toolInfo := scanForCppCheck()
	assert.True(t, toolInfo.Available)
	assert.Equal(t, "1.86", toolInfo.Version)
}

func TestScanForCppCheckExecutionError(t *testing.T) {
	expectedError := errors.New("Execution error")
	runner = runnerDummy{
		outputCppCheck: nil,
		errCppCheck:    expectedError,
	}
	defer func() { runner = toolrunner.NewCommandlineRunner() }()
	toolInfo := scanForCppCheck()
	assert.False(t, toolInfo.Available)
	assert.Equal(t, "", toolInfo.Version)
}

func TestScanForCppCheckInvalidToolOutput(t *testing.T) {
	runner = runnerDummy{
		outputCppCheck: []byte(strings.Trim(actualCppCheckResult, "\n")),
		errCppCheck:    nil,
	}
	defer func() { runner = toolrunner.NewCommandlineRunner() }()
	toolInfo := scanForCppCheck()
	assert.False(t, toolInfo.Available)
	assert.Equal(t, "", toolInfo.Version)
}

//=============================================================================
func TestScanTools(t *testing.T) {
	runner = runnerDummy{
		outputCppCheck: []byte(actualCppCheckResult),
		errCppCheck:    nil,
		outputCmake:    []byte(actualCmakeResult),
		errCmake:       nil,
	}

	infos, err := ScanTools()
	assert.Equal(t, nil, err)
	assert.True(t, infos.CMake.Available)
	assert.True(t, infos.CppCheck.Available)
	assert.Equal(t, "3.14.3", infos.CMake.Version)
	assert.Equal(t, "1.86", infos.CppCheck.Version)
}
