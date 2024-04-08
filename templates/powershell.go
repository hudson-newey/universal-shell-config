package templates

import "usc/config"

func PowershellAliases(alias *config.Alias) string {
	var returnTemplate string

	for _, shortHand := range alias.Aliases {
		template := "function " + shortHand + "-func {\n" +
			"\t" + alias.Command + "\n}\n" +
			"Set-Alias " + shortHand + " " + shortHand + "-func\n"
		returnTemplate += template
	}

	return returnTemplate
}

func PowershellSource(source *config.Source) string {
	template := ". " + source.Path + "\n"
	return template
}
