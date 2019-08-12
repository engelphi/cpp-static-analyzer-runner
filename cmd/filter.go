package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/engelphi/cpp-static-analyzer-runner/ccdb"
	"github.com/spf13/cobra"
)

var ccdbFile string
var filterPrefix string

func init() {
	rootCmd.AddCommand(filterCmd)
	filterCmd.Flags().StringVarP(&ccdbFile, "ccdb-file", "", "", "path to your compile_commands.json")
	filterCmd.Flags().StringVarP(&filterPrefix, "path-prefix", "p", "", "Path prefix that should be filtered")
	filterCmd.MarkFlagRequired("ccdb-file")
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

var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filters commands from a given compile_commands.json",
	Long: `Filters commands from a given compile_commands.json
					in order to avoid analyzing files in which you are not
					interested`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Filter: " + strings.Join(args, ""))
		log.Println("file: " + ccdbFile)
		log.Println("prefix: " + filterPrefix)

		if !validator.IsValidDirectory(filterPrefix) {
			fmt.Println("Invalid directory path")
			os.Exit(1)
		}

		commands := ccdb.NewCompileCommandsDB()

		file, err := os.Open(ccdbFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		err = commands.LoadCompileCommands(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = commands.Filter(filterPrefix)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		outputFile, err := os.Create(ccdbFile + ".filtered")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer outputFile.Close()

		err = commands.WriteCompileCommands(outputFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		outputFile.Sync()
	},
}
