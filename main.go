package main

import (
	"fmt"
	"log"

	"github.com/engelphi/cpp-static-analyzer-runner/toolrunner"
)

func main() {
	runner := toolrunner.NewCommandlineRunner()
	_, err := runner.RunCmake("-S", ".", "-B", "build", "-DCMAKE_EXPORT_COMPILE_COMMANDS=ON")
	if err != nil {
		log.Fatal(err)
	}

	out, err := runner.RunCppCheck("--project=build/compile_commands.json", "--enable=all", "-q", "--xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
