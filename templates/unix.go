package templates

import (
	"usc/config"
)

func UnixAliases(alias *config.Alias) string {
	var returnTemplate string

	for _, shortHand := range alias.Aliases {
		template := "alias " + shortHand + "='" + alias.Command + "'\n"
		returnTemplate += template
	}

	return returnTemplate
}

func UnixSource(source *config.Source) string {
	template := "source " + source.Path + "\n"
	return template
}
