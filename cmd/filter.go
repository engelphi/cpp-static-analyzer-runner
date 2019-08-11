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

		commands := ccdb.NewCompileCommandsDB()

		err := commands.LoadCompileCommands(ccdbFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = commands.Filter(filterPrefix)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = commands.WriteCompileCommands(ccdbFile + ".filtered")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
