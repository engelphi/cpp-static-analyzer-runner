package ccdb

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fileValidatorMock struct {
	ExpectedValidFileReturn      bool
	ExpectedValidDirectoryReturn bool
}

func (f fileValidatorMock) IsValidFile(filename string) bool {
	return f.ExpectedValidFileReturn
}

func (f fileValidatorMock) IsValidDirectory(directoryPath string) bool {
	return f.ExpectedValidDirectoryReturn
}

// =================================================================================================
var testCommands = []byte(`[
	{
		"directory": "/home/pen/Project/algorithm/build",
		"command": "/usr/bin/c++   -I/home/pen/Project/algorithm/algorithms -I/home/pen/Project/algorithm/catch   -Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -Wold-style-cast -Wcast-align -Wunused -Woverloaded-virtual -Wpedantic -Wconversion -Wsign-conversion -Wnull-dereference -Wdouble-promotion -Wformat=2   -std=c++17 -o CMakeFiles/test-algorithm.dir/test/intersperse.cpp.o -c /home/pen/Project/algorithm/test/intersperse.cpp",
		"file": "/home/pen/Project/algorithm/test/intersperse.cpp"
	},
	{
		"directory": "/home/pen/Project/algorithm/build",
		"command": "/usr/bin/c++   -I/home/pen/Project/algorithm/algorithms -I/home/pen/Project/algorithm/catch   -Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -Wold-style-cast -Wcast-align -Wunused -Woverloaded-virtual -Wpedantic -Wconversion -Wsign-conversion -Wnull-dereference -Wdouble-promotion -Wformat=2   -std=c++17 -o CMakeFiles/test-algorithm.dir/test/main.cpp.o -c /home/pen/Project/algorithm/test/main.cpp",
		"file": "/home/pen/Project/algorithm/test/main.cpp"
	}
]
`)

// =================================================================================================
func TestParseCompileCommands(t *testing.T) {
	testCommandReader := bytes.NewReader(testCommands)

	expected := []CompileCommand{
		{
			Directory: "/home/pen/Project/algorithm/build",
			Command:   "/usr/bin/c++   -I/home/pen/Project/algorithm/algorithms -I/home/pen/Project/algorithm/catch   -Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -Wold-style-cast -Wcast-align -Wunused -Woverloaded-virtual -Wpedantic -Wconversion -Wsign-conversion -Wnull-dereference -Wdouble-promotion -Wformat=2   -std=c++17 -o CMakeFiles/test-algorithm.dir/test/intersperse.cpp.o -c /home/pen/Project/algorithm/test/intersperse.cpp",
			File:      "/home/pen/Project/algorithm/test/intersperse.cpp",
		},
		{
			Directory: "/home/pen/Project/algorithm/build",
			Command:   "/usr/bin/c++   -I/home/pen/Project/algorithm/algorithms -I/home/pen/Project/algorithm/catch   -Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -Wold-style-cast -Wcast-align -Wunused -Woverloaded-virtual -Wpedantic -Wconversion -Wsign-conversion -Wnull-dereference -Wdouble-promotion -Wformat=2   -std=c++17 -o CMakeFiles/test-algorithm.dir/test/main.cpp.o -c /home/pen/Project/algorithm/test/main.cpp",
			File:      "/home/pen/Project/algorithm/test/main.cpp",
		},
	}

	commands, err := parseCompileCommands(testCommandReader)
	assert.Nil(t, err)
	assert.Equal(t, expected, commands)
}

func TestParseCompileCommandsFailsIfEmpty(t *testing.T) {
	testCommandReader := bytes.NewReader([]byte(""))
	_, err := parseCompileCommands(testCommandReader)
	assert.NotNil(t, err)
}

// =================================================================================================

func TestFilter(t *testing.T) {
	validator = fileValidatorMock{
		ExpectedValidDirectoryReturn: true,
		ExpectedValidFileReturn:      false,
	}
	defer func() { validator = fileValidatorImpl{} }()
	input := CompileCommands{
		Commands: []CompileCommand{
			{
				Directory: "/home/pen/Project/algorithm/build",
				Command:   "/usr/bin/c++   -I/home/pen/Project/algorithm/algorithms -I/home/pen/Project/algorithm/catch   -Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -Wold-style-cast -Wcast-align -Wunused -Woverloaded-virtual -Wpedantic -Wconversion -Wsign-conversion -Wnull-dereference -Wdouble-promotion -Wformat=2   -std=c++17 -o CMakeFiles/test-algorithm.dir/test/intersperse.cpp.o -c /home/pen/Project/algorithm/test/intersperse.cpp",
				File:      "/home/pen/Project/algorithm/test/intersperse.cpp",
			},
			{
				Directory: "/home/pen/Project/algorithm/build",
				Command:   "/usr/bin/c++   -I/home/pen/Project/algorithm/algorithms -I/home/pen/Project/algorithm/catch   -Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -Wold-style-cast -Wcast-align -Wunused -Woverloaded-virtual -Wpedantic -Wconversion -Wsign-conversion -Wnull-dereference -Wdouble-promotion -Wformat=2   -std=c++17 -o CMakeFiles/test-algorithm.dir/test/main.cpp.o -c /home/pen/Project/algorithm/test/main.cpp",
				File:      "/home/pen/Project/algorithm/test2/main.cpp",
			},
		},
	}

	expected := CompileCommands{
		Commands: []CompileCommand{
			{
				Directory: "/home/pen/Project/algorithm/build",
				Command:   "/usr/bin/c++   -I/home/pen/Project/algorithm/algorithms -I/home/pen/Project/algorithm/catch   -Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -Wold-style-cast -Wcast-align -Wunused -Woverloaded-virtual -Wpedantic -Wconversion -Wsign-conversion -Wnull-dereference -Wdouble-promotion -Wformat=2   -std=c++17 -o CMakeFiles/test-algorithm.dir/test/intersperse.cpp.o -c /home/pen/Project/algorithm/test/intersperse.cpp",
				File:      "/home/pen/Project/algorithm/test/intersperse.cpp",
			},
		},
	}

	err := input.Filter("/home/pen/Project/algorithm/test2/")
	assert.Nil(t, err)
	assert.Equal(t, expected, input)
}

func TestFilterFailsIfInvalidDirectoryIsGiven(t *testing.T) {
	validator = fileValidatorMock{
		ExpectedValidDirectoryReturn: false,
		ExpectedValidFileReturn:      false,
	}
	defer func() { validator = fileValidatorImpl{} }()
	input := CompileCommands{
		Commands: []CompileCommand{
			{
				Directory: "/home/pen/Project/algorithm/build",
				Command:   "/usr/bin/c++   -I/home/pen/Project/algorithm/algorithms -I/home/pen/Project/algorithm/catch   -Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -Wold-style-cast -Wcast-align -Wunused -Woverloaded-virtual -Wpedantic -Wconversion -Wsign-conversion -Wnull-dereference -Wdouble-promotion -Wformat=2   -std=c++17 -o CMakeFiles/test-algorithm.dir/test/intersperse.cpp.o -c /home/pen/Project/algorithm/test/intersperse.cpp",
				File:      "/home/pen/Project/algorithm/test/intersperse.cpp",
			},
			{
				Directory: "/home/pen/Project/algorithm/build",
				Command:   "/usr/bin/c++   -I/home/pen/Project/algorithm/algorithms -I/home/pen/Project/algorithm/catch   -Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -Wold-style-cast -Wcast-align -Wunused -Woverloaded-virtual -Wpedantic -Wconversion -Wsign-conversion -Wnull-dereference -Wdouble-promotion -Wformat=2   -std=c++17 -o CMakeFiles/test-algorithm.dir/test/main.cpp.o -c /home/pen/Project/algorithm/test/main.cpp",
				File:      "/home/pen/Project/algorithm/test2/main.cpp",
			},
		},
	}

	err := input.Filter("/home/pen/Project/algorithm/test2/")
	assert.NotNil(t, err)
}
