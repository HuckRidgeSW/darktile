package config

import "github.com/liamg/darktile/internal/app/darktile/termutil"

func GetDefaultTheme() Theme { return defaultTheme }

func LoadThemeFromConf(conf *Config, themeConf *Theme) (*termutil.Theme, error) {
	return loadThemeFromConf(conf, themeConf)
}
