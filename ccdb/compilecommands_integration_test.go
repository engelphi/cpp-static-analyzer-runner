// +build integration

package ccdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/udhos/equalfile"
)

func TestLoadCompileCommands(t *testing.T) {
	ccdb := NewCompileCommandsDB()
	var path string = "./testdata/compile_commands.json"
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

	err := ccdb.LoadCompileCommands(path)
	assert.Nil(t, err)
	assert.Equal(t, &expected, ccdb)
}

func TestLoadCompileCommandsFailsWithNonExistingFile(t *testing.T) {
	ccdb := NewCompileCommandsDB()
	var path string = "./notexistent.json"
	err := ccdb.LoadCompileCommands(path)
	assert.NotNil(t, err)
}

func TestLoadCompileCommandsFailsWhenMalformedContentIsDiscovered(t *testing.T) {
	ccdb := NewCompileCommandsDB()
	var path string = "./testdata/compile_commands.json.malformed"
	err := ccdb.LoadCompileCommands(path)
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

	err := input.WriteCompileCommands("./testdata/testoutput.json")
	assert.Nil(t, err)

	cmp := equalfile.New(nil, equalfile.Options{}) // compare using single mode
	equal, err := cmp.CompareFile("./testdata/testoutput.json", "./testdata/compile_commands.json")
	assert.Nil(t, err)
	assert.True(t, equal)

	err = os.Remove("./testdata/testoutput.json")
	assert.Nil(t, err)
}
