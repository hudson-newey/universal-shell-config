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

func PowershellSource(source string) string {
	return ". " + source + "\n"
}

func PowershellPath(path string) string {
	return "$env:PATH += \";" + path + "\"\n"
}

func PowershellEnvVariable(variable *config.EnvVariable) string {
	return "$env:" + variable.Name + " = \"" + variable.Value + "\"\n"
}
