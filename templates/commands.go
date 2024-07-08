package templates

func BuildCommand(command string) string {
	if command == "" {
		return ""
	}

	return command + "\n"
}
