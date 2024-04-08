package build

import (
	"usc/config"
	"usc/templates"
)

type BuildConfig struct {
	AliasConverter   func(alias *config.Alias) string
	SourceConverter  func(source string) string
	CommandConverter func(command string) string
	PathConverter    func(path string) string
	outFile          string
}

var defaultBuildConfigs = []BuildConfig{
	{
		AliasConverter:   templates.UnixAliases,
		SourceConverter:  templates.UnixSource,
		CommandConverter: templates.BuildCommand,
		PathConverter:    templates.UnixPath,
		outFile:          ".zshrc",
	},
	{
		AliasConverter:   templates.UnixAliases,
		SourceConverter:  templates.UnixSource,
		CommandConverter: templates.BuildCommand,
		PathConverter:    templates.UnixPath,
		outFile:          ".bashrc",
	},
	{
		AliasConverter:   templates.PowershellAliases,
		SourceConverter:  templates.PowershellSource,
		CommandConverter: templates.BuildCommand,
		PathConverter:    templates.PowershellPath,
		outFile:          "Profile.ps1",
	},
}
