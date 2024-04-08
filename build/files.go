package build

import "os"

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func deleteFile(path string) error {
	return os.Remove(path)
}

func appendToFile(path string, contents string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(contents); err != nil {
		panic(err)
	}
}
