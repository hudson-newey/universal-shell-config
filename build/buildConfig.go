package build

import (
	"usc/config"
	"usc/templates"
	"usc/util"
)

type BuildConfig struct {
	AliasConverter   func(alias *config.Alias) string
	SourceConverter  func(source string) string
	CommandConverter func(command string) string
	PathConverter    func(path string) string
	OutFile          string
	InstallPath      string
}

const WILDCARD_TARGET string = "*"

var defaultBuildConfigs = []BuildConfig{
	{
		AliasConverter:   templates.UnixAliases,
		SourceConverter:  templates.UnixSource,
		CommandConverter: templates.BuildCommand,
		PathConverter:    templates.UnixPath,
		OutFile:          ".zshrc",
		InstallPath:      ".zshrc",
	},
	{
		AliasConverter:   templates.UnixAliases,
		SourceConverter:  templates.UnixSource,
		CommandConverter: templates.BuildCommand,
		PathConverter:    templates.UnixPath,
		OutFile:          ".bashrc",
		InstallPath:      ".bashrc",
	},
	// TODO: I've hard coded the install path because I have to call $PROFILE inside a powershell session
	// this is obviously not ideal. maybe a set of install scripts for every build config?
	{
		AliasConverter:   templates.PowershellAliases,
		SourceConverter:  templates.PowershellSource,
		CommandConverter: templates.BuildCommand,
		PathConverter:    templates.PowershellPath,
		OutFile:          "Profile.ps1",
		InstallPath:      "Profile.ps1",
	},
}

func OutputSatisfiesTarget(outFile string, target []string) bool {
	if util.ValueInArray[string](WILDCARD_TARGET, target) {
		return true
	}

	for _, t := range target {
		if t == outFile {
			return true
		}
	}

	return false
}
