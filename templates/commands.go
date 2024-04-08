package templates

import "usc/config"

func BuildCommand(command *config.Command) string {
	return command.Command + "\n"
}
