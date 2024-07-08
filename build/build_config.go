package build

import (
	"strings"
	"usc/config"
	"usc/templates"
	"usc/util"
)

type BuildConfig struct {
	EnvVariableConverter func(alias *config.EnvVariable) string
	AliasConverter       func(alias *config.Alias) string
	SourceConverter      func(source string) string
	CommandConverter     func(command string) string
	PathConverter        func(path string) string
	OutFile              string
	InstallPath          string
}

const WILDCARD_TARGET string = "*"

var defaultBuildConfigs = []BuildConfig{
	{
		EnvVariableConverter: templates.UnixEnvVariable,
		AliasConverter:       templates.UnixAliases,
		SourceConverter:      templates.UnixSource,
		CommandConverter:     templates.BuildCommand,
		PathConverter:        templates.UnixPath,
		OutFile:              ".zshrc",
		InstallPath:          ".zshrc",
	},
	{
		EnvVariableConverter: templates.UnixEnvVariable,
		AliasConverter:       templates.UnixAliases,
		SourceConverter:      templates.UnixSource,
		CommandConverter:     templates.BuildCommand,
		PathConverter:        templates.UnixPath,
		OutFile:              ".bashrc",
		InstallPath:          ".bashrc",
	},
	// TODO: I've hard coded the install path because I have to call $PROFILE inside a powershell session
	// this is obviously not ideal. maybe a set of install scripts for every build config?
	{
		EnvVariableConverter: templates.PowershellEnvVariable,
		AliasConverter:       templates.PowershellAliases,
		SourceConverter:      templates.PowershellSource,
		CommandConverter:     templates.BuildCommand,
		PathConverter:        templates.PowershellPath,
		OutFile:              "Profile.ps1",
		InstallPath:          "Profile.ps1",
	},
}

func OutputSatisfiesTarget(outPath string, targets []string) bool {
	if util.ValueInArray(WILDCARD_TARGET, targets) {
		return true
	}

	outPathSplit := strings.Split(outPath, "/")
	outFile := outPathSplit[len(outPathSplit)-1]

	return util.ValueInArray(outFile, targets)
}
