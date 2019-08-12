package ccdb

import (
	"encoding/json"
	"io"
	"log"
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
	LoadCompileCommands(data io.Reader) error
	Filter(directory string) error
	WriteCompileCommands(outputBuffer io.Writer) error
}

// CompileCommands Represents the contents of a compile_commands.json
type CompileCommands struct {
	Commands []CompileCommand
}

// NewCompileCommandsDB Creates a new CompileCommandsDB
func NewCompileCommandsDB() CompileCommandsDB {
	return &CompileCommands{}
}

// Filter filters out commands on files that start are situated in the given directory
func (c *CompileCommands) Filter(directory string) error {
	var filteredCommands []CompileCommand

	for _, command := range c.Commands {
		if !strings.HasPrefix(command.File, directory) {
			filteredCommands = append(filteredCommands, command)
		}
	}

	c.Commands = filteredCommands
	return nil
}

// LoadCompileCommands loads compile commands from a given file. Returns an error if it fails to load the commands
func (c *CompileCommands) LoadCompileCommands(data io.Reader) error {
	dec := json.NewDecoder(data)
	if err := dec.Decode(&c.Commands); err != nil {
		log.Println("Failed to parse compile commands")
		return err
	}

	return nil
}

// WriteCompileCommands writes the compile commands database
func (c *CompileCommands) WriteCompileCommands(outputBuffer io.Writer) error {
	enc := json.NewEncoder(outputBuffer)
	enc.SetIndent("", "\t")
	if err := enc.Encode(&c.Commands); err != nil {
		log.Println("Failed to serialize compile commands")
		return err
	}
	return nil
}
