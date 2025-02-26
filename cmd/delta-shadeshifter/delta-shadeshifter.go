package deltashadeshifter

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/tomasaschan/delta-shadeshifter/pkg/darkmode"
)

func Run(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) error {
	if dark, err := darkmode.IsDarkMode(); err == nil && dark {
		os.Setenv("BAT_THEME", "Visual Studio Dark+")
	} else if err == nil {
		os.Setenv("BAT_THEME", "GitHub")
	} else {
		return fmt.Errorf("detect dark mode: %w", err)
	}

	cmd := exec.Command("delta", args...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return cmd.Run()
}
