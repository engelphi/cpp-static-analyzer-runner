package ccdb

import (
	"encoding/json"
	"fmt"
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

type fileValidator interface {
	IsValidFile(filename string) bool
	IsValidDirectory(directoryPath string) bool
}

type fileValidatorImpl struct {
}

func (f fileValidatorImpl) IsValidFile(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	if mode := fileInfo.Mode(); !mode.IsRegular() {
		return false
	}

	return true
}

func (f fileValidatorImpl) IsValidDirectory(directoryPath string) bool {
	fileInfo, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		return false
	}

	if mode := fileInfo.Mode(); !mode.IsDir() {
		return false
	}

	return true
}

var validator fileValidator = fileValidatorImpl{}

// CompileCommandsDB interface for interacting with compile_commands.json
type CompileCommandsDB interface {
	Filter(directory string) (CompileCommandsDB, error)
}

// CompileCommands Represents the contents of a compile_commands.json
type CompileCommands struct {
	Commands []CompileCommand
}

// Filter filters out commands on files that start are situated in the given directory
func (c CompileCommands) Filter(directory string) (CompileCommandsDB, error) {
	if !validator.IsValidDirectory(directory) {
		return CompileCommands{}, fmt.Errorf("Invalid directory path")
	}

	var filteredCommands CompileCommands

	for _, command := range c.Commands {
		if !strings.HasPrefix(command.File, directory) {
			filteredCommands.Commands = append(filteredCommands.Commands, command)
		}
	}

	return filteredCommands, nil
}

// LoadCompileCommands loads compile commands from a given file. Returns an error if it fails to load the commands
func LoadCompileCommands(filename string) (CompileCommandsDB, error) {
	if !validator.IsValidFile(filename) {
		return CompileCommands{}, fmt.Errorf("Invalid path to compile commands database")
	}

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
