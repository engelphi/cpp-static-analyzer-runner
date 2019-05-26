package compilecommandsdb

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	expected := CompileCommands{
		Commands: []CompileCommand{
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
		},
	}

	commands, err := ParseCompileCommands(testCommandReader)
	assert.Nil(t, err)
	assert.Equal(t, expected, commands)
}

func TestParseCompileCommandsFailsIfEmpty(t *testing.T) {
	testCommandReader := bytes.NewReader([]byte(""))
	_, err := ParseCompileCommands(testCommandReader)
	assert.NotNil(t, err)
}

// =================================================================================================
func TestWriteCompileCommands(t *testing.T) {
	expected := testCommands

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
				File:      "/home/pen/Project/algorithm/test/main.cpp",
			},
		},
	}

	buf := []byte("")
	outputBuffer := bytes.NewBuffer(buf)
	err := WriteCompileCommands(outputBuffer, input)
	assert.Nil(t, err)
	assert.Equal(t, string(expected), string(outputBuffer.Bytes()))
}

// =================================================================================================

func TestFilter(t *testing.T) {
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

	actual := input.Filter("/home/pen/Project/algorithm/test2/")
	assert.Equal(t, expected, actual)
}
