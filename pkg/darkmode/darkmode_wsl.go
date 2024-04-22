//go:build linux

package darkmode

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func IsDarkMode() (bool, error) {
	if f, err := os.OpenFile("/proc/version", os.O_RDONLY, 0); err != nil {
		return false, fmt.Errorf("open /proc/version: %w", err)
	} else if data, err := io.ReadAll(f); err != nil {
		return false, fmt.Errorf("read /proc/version: %w", err)
	} else if bytes.Contains(data, []byte("WSL")) {
		return isDarkModeWSL()
	}

	return false, errors.New("unsupported platform")
}

func isDarkModeWSL() (bool, error) {
	cmd := exec.Command("powershell.exe", `Get-ItemProperty -Path Registry::HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Themes\Personalize\ -name SystemUsesLightTheme | Select-Object -ExpandProperty SystemUsesLightTheme`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return false, fmt.Errorf("connect to stdout: %w", err)
	}

	buf := make([]byte, 1)
	go func() {
		if _, err := stdout.Read(buf); err != nil && err != io.EOF {
			panic(fmt.Errorf("read stdout: %w", err))
		}
	}()

	if err := cmd.Run(); err != nil {
		return false, fmt.Errorf("run command: %w", err)
	}

	isDarkMode := buf[0] == byte('0')
	return isDarkMode, nil
}
