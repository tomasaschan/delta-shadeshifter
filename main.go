package main

import (
	"fmt"
	"os"
	"os/exec"

	deltashadeshifter "github.com/tomasaschan/delta-shadeshifter/cmd/delta-shadeshifter"
)

func main() {
	if err := deltashadeshifter.Run(os.Stdin, os.Stdout, os.Stderr, os.Args[1:]); err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			os.Exit(e.ExitCode())
		}

		fmt.Fprintf(os.Stderr, "delta-shadeshifter: %v\n", err)
		os.Exit(1)
	}
}
