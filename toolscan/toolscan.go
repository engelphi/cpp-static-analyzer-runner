package toolscan

import (
	"bytes"
	"strings"

	"github.com/engelphi/cpp-static-analyzer-runner/toolrunner"
)

var runner = toolrunner.NewCommandlineRunner()

// ToolInfo contains information about a tool
type ToolInfo struct {
	Available bool
	Version   string
}

// ScanResult contains the result of a tool scan
type ScanResult struct {
	CMake    ToolInfo
	CppCheck ToolInfo
}

func scanForCmake() (info ToolInfo) {
	out, err := runner.RunCmake("--version")
	if err != nil {
		info.Available = false
		info.Version = ""
		return
	}

	var buf = bytes.NewBuffer(out)
	line, err := buf.ReadString('\n')
	if err != nil {
		info.Available = false
		info.Version = ""
		return
	}

	words := strings.Split(line, " ")
	info.Available = true
	info.Version = strings.Trim(words[len(words)-1], "\n")
	return
}

func scanForCppCheck() (info ToolInfo) {
	out, err := runner.RunCppCheck("--version")
	if err != nil {
		info.Available = false
		info.Version = ""
		return
	}

	var buf = bytes.NewBuffer(out)
	line, err := buf.ReadString('\n')
	if err != nil {
		info.Available = false
		info.Version = ""
		return
	}

	words := strings.Split(line, " ")
	info.Available = true
	info.Version = strings.Trim(words[len(words)-1], "\n")
	return
}

// ScanTools scans for the available tools
func ScanTools() (tools ScanResult, err error) {
	tools.CMake = scanForCmake()
	tools.CppCheck = scanForCppCheck()
	return tools, nil
}
