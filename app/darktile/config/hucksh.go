package config

import (
	"github.com/liamg/darktile/app/darktile/termutil"
	p "github.com/liamg/darktile/internal/app/darktile/config"
)

type (
	Config = p.Config
	Theme  = p.Theme
)

func LoadConfig() (*Config, error)                       { return p.LoadConfig() }
func DefaultConfig() *Config                             { return p.DefaultConfig() }
func DefaultTheme(conf *Config) (*termutil.Theme, error) { return p.DefaultTheme(conf) }
func GetDefaultTheme() Theme                             { return p.GetDefaultTheme() }

func LoadThemeFromConf(conf *Config, themeConf *Theme) (*termutil.Theme, error) {
	return p.LoadThemeFromConf(conf, themeConf)
}
