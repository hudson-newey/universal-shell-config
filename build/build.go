package build

import (
	"fmt"
	"usc/config"
)

func buildWarnings(filePath string) {
	warnings := "# DO NOT EDIT. Configuration generated by usc\n"
	appendToFile(filePath, warnings)
}

func BuildTemplates(configuration *config.Config, outDir string) {
	var derivedBuildConfigs []BuildConfig

	for _, buildConfig := range defaultBuildConfigs {
		derivedBuildConfig := buildConfig
		derivedBuildConfig.outFile = outDir + "/" + derivedBuildConfig.outFile
		derivedBuildConfigs = append(derivedBuildConfigs, derivedBuildConfig)
	}

	for _, bc := range derivedBuildConfigs {
		if fileExists(bc.outFile) {
			deleteFile(bc.outFile)
			buildWarnings(bc.outFile)
		}
	}

	fmt.Println()
	for _, buildConfig := range derivedBuildConfigs {
		fmt.Println("Building " + buildConfig.outFile)

		for _, alias := range configuration.Aliases {
			aliasTemplate := buildConfig.AliasConverter(alias)
			appendToFile(buildConfig.outFile, aliasTemplate)
		}

		for _, source := range configuration.Sources {
			sourceTemplate := buildConfig.SourceConverter(source)
			appendToFile(buildConfig.outFile, sourceTemplate)
		}

		for _, path := range configuration.Paths {
			pathTemplate := buildConfig.PathConverter(path)
			appendToFile(buildConfig.outFile, pathTemplate)
		}

		// pre-interaction commands should always come last because they sometimes depend on
		// aliases, sources, and paths
		for _, command := range configuration.Commands {
			commandTemplate := buildConfig.CommandConverter(command)
			appendToFile(buildConfig.outFile, commandTemplate)
		}
	}

	fmt.Println("\nDone!")
}
