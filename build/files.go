package build

import (
	"io"
	"os"
)

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

func createDirectory(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func directoryExists(path string) bool {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return stat.IsDir()

}

func copyFile(inPath string, outPath string) {
	inFile, err := os.Open(inPath)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		panic(err)
	}
}
