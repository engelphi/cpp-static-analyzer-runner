package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cppsar",
	Short: "cppsar is a runner for C++ static analysis tools",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello")
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
