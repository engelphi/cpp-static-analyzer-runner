// +build integration

package compilecommandsdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadCompileCommands(t *testing.T) {
	var path string = "./compile_commands.json"
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

	actual, err := LoadCompileCommands(path)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestLoadCompileCommandsFailsWithNonExistingFile(t *testing.T) {
	var path string = "./test.json"
	_, err := LoadCompileCommands(path)
	assert.NotNil(t, err)
}

func TestLoadCompileCommandsFailsWhenMalformedContentIsDiscovered(t *testing.T) {
	var path string = "./compile_commands.json.malformed"
	_, err := LoadCompileCommands(path)
	assert.NotNil(t, err)
}
