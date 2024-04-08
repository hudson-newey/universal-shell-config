package build

import (
	"usc/config"
	"usc/templates"
)

type BuildConfig struct {
	AliasConverter   func(alias *config.Alias) string
	SourceConverter  func(source *config.Source) string
	CommandConverter func(command *config.Command) string
	outFile          string
}

var defaultBuildConfigs = []BuildConfig{
	BuildConfig{
		AliasConverter:   templates.UnixAliases,
		SourceConverter:  templates.UnixSource,
		CommandConverter: templates.BuildCommand,
		outFile:          ".zshrc",
	},
	BuildConfig{
		AliasConverter:   templates.UnixAliases,
		SourceConverter:  templates.UnixSource,
		CommandConverter: templates.BuildCommand,
		outFile:          ".bashrc",
	},
	BuildConfig{
		AliasConverter:   templates.UnixAliases,
		SourceConverter:  templates.UnixSource,
		CommandConverter: templates.BuildCommand,
		outFile:          ".profile",
	},
	BuildConfig{
		AliasConverter:   templates.PowershellAliases,
		SourceConverter:  templates.PowershellSource,
		CommandConverter: templates.BuildCommand,
		outFile:          "Profile.ps1",
	},
}
