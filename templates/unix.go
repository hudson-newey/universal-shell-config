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

func UnixSource(source string) string {
	return "source " + source + "\n"
}

func UnixPath(path string) string {
	return "export PATH=$PATH:" + path + "\n"
}
