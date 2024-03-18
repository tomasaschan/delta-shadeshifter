package darkmode

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

const (
	RegistryKey          = `Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`
	SystemUsesLightTheme = `SystemUsesLightTheme`
	AppsUseLightTheme    = `AppsUseLightTheme`
)

func IsDarkMode() (bool, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, RegistryKey, registry.QUERY_VALUE)
	if err != nil {
		return false, fmt.Errorf("open key %s: %w", SystemUsesLightTheme, err)
	}

	buf := make([]byte, registry.DWORD)

	if _, _, err := k.GetValue(AppsUseLightTheme, buf); err != nil {
		return false, fmt.Errorf("get string value %s: %w", SystemUsesLightTheme, err)
	}
	return buf[0] != 0x1, nil
}
