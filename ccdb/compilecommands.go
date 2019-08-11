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
	LoadCompileCommands(filename string) error
	Filter(directory string) error
	WriteCompileCommands(filename string) error
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
	if !validator.IsValidDirectory(directory) {
		return fmt.Errorf("Invalid directory path")
	}

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
func (c *CompileCommands) LoadCompileCommands(filename string) error {
	if !validator.IsValidFile(filename) {
		return fmt.Errorf("Invalid path to compile commands database")
	}

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Println("Failed to open compile commands file")
		return err
	}

	commands, err := parseCompileCommands(file)
	if err != nil {
		log.Println("Failed to load compile commands")
		return err
	}

	c.Commands = commands
	return nil
}

// ParseCompileCommands parses compile commands from the given reader
func parseCompileCommands(r io.Reader) ([]CompileCommand, error) {
	var compileCommands []CompileCommand
	dec := json.NewDecoder(r)

	if err := dec.Decode(&compileCommands); err != nil {
		log.Println("Failed to parse compile commands")
		return []CompileCommand{}, err
	}

	return compileCommands, nil
}

// WriteCompileCommands writes the compile commands database
func (c *CompileCommands) WriteCompileCommands(filename string) error {
	outputFile, err := os.Create(filename)
	if err != nil {
		log.Println("Failed to create output file")
		return err
	}
	defer outputFile.Close()

	enc := json.NewEncoder(outputFile)
	enc.SetIndent("", "\t")

	if err := enc.Encode(&c.Commands); err != nil {
		log.Println("Failed to serialize compile commands")
		return err
	}

	outputFile.Sync()

	return nil
}
