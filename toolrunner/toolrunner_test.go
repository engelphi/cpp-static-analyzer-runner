package toolrunner

import (
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
	runner := NewCommandlineRunner()
	out, err := runner.RunCmake("--version")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedCmakeResult, string(out))
}

func TestRunCppCheck(t *testing.T) {
	const expectedCppCheckResult = "Cppcheck 1.86\n"
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	runner := NewCommandlineRunner()
	out, err := runner.RunCppCheck("--version")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedCppCheckResult, string(out))
}
