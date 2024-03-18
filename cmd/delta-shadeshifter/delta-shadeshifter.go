package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/tomasaschan/delta-shadeshifter/pkg/darkmode"
)

func main() {
	if err := run(os.Stdin, os.Stdout, os.Stderr, os.Args[1:]); err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			os.Exit(e.ExitCode())
		}

		fmt.Fprintf(os.Stderr, "delta-shadeshifter: %v\n", err)
		os.Exit(1)
	}
}

func run(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) error {
	if dark, err := darkmode.IsDarkMode(); err == nil && dark {
		os.Setenv("BAT_THEME", "Visual Studio Dark+")
	} else if err == nil {
		os.Setenv("BAT_THEME", "Visual Studio Light+")
	} else {
		return fmt.Errorf("detect dark mode: %w", err)
	}

	cmd := exec.Command("delta", args...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return cmd.Run()
}
