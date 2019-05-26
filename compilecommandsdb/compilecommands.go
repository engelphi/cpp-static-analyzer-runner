package compilecommandsdb

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

// CompileCommand Represents a single compile command
type CompileCommand struct {
	Directory string `json:"directory"`
	Command   string `json:"command"`
	File      string `json:"file"`
}

// CompileCommandsDB interface for interacting with compile_commands.json
type CompileCommandsDB interface {
	Filter(directory string) CompileCommandsDB
}

// CompileCommands Represents the contents of a compile_commands.json
type CompileCommands struct {
	Commands []CompileCommand
}

// Filter filters out commands on files that start are situated in the given directory
func (c CompileCommands) Filter(directory string) CompileCommandsDB {
	var filteredCommands CompileCommands

	for _, command := range c.Commands {
		if !strings.HasPrefix(command.File, directory) {
			filteredCommands.Commands = append(filteredCommands.Commands, command)
		}
	}

	return filteredCommands
}

// LoadCompileCommands loads compile commands from a given file. Returns an error if it fails to load the commands
func LoadCompileCommands(filename string) (CompileCommandsDB, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Println("Failed to open compile commands file")
		return CompileCommands{}, err
	}

	commands, err := ParseCompileCommands(file)
	if err != nil {
		log.Println("Failed to load compile commands")
		return CompileCommands{}, err
	}

	return commands, nil
}

// ParseCompileCommands parses compile commands from the given reader
func ParseCompileCommands(r io.Reader) (CompileCommands, error) {
	var compileCommands CompileCommands
	dec := json.NewDecoder(r)

	if err := dec.Decode(&compileCommands.Commands); err != nil {
		log.Println("Failed to parse compile commands")
		return CompileCommands{}, err
	}

	return compileCommands, nil
}

// WriteCompileCommands writes the compile commands database
func WriteCompileCommands(w io.Writer, commands CompileCommands) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")

	if err := enc.Encode(&commands.Commands); err != nil {
		log.Println("Failed to serialize compile commands")
		return err
	}

	return nil
}
