package ccdb

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// =================================================================================================
var testCommands = string(`[
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

var testCommandsMalformed = string(`[
	{
		"directory": "/home/pen/Project/algorithm/build",
		"command": "/usr/bin/c++   -I/home/pen/Project/algorithm/algorithms -I/home/pen/Project/algorithm/catch   -Wall -Wextra -Wshadow -Wnon-virtual-dtor -pedantic -Wold-style-cast -Wcast-align -Wunused -Woverloaded-virtual -Wpedantic -Wconversion -Wsign-conversion -Wnull-dereference -Wdouble-promotion -Wformat=2   -std=c++17 -o CMakeFiles/test-algorithm.dir/test/intersperse.cpp.o -c /home/pen/Project/algorithm/test/intersperse.cpp",
		"fil
]`)

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

	err := input.Filter("/home/pen/Project/algorithm/test2/")
	assert.Nil(t, err)
	assert.Equal(t, expected, input)
}

func TestLoadCompileCommands(t *testing.T) {
	ccdb := NewCompileCommandsDB()
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

	err := ccdb.LoadCompileCommands(strings.NewReader(testCommands))
	assert.Nil(t, err)
	assert.Equal(t, &expected, ccdb)
}

func TestLoadCompileCommandsFailsWhenMalformedContentIsDiscovered(t *testing.T) {
	ccdb := NewCompileCommandsDB()
	err := ccdb.LoadCompileCommands(strings.NewReader(testCommandsMalformed))
	assert.NotNil(t, err)
}

func TestWriteCompileCommands(t *testing.T) {
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

	var buf bytes.Buffer
	err := input.WriteCompileCommands(&buf)
	assert.Nil(t, err)
	assert.Equal(t, testCommands, buf.String())
}
