package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"usc/build"
	"usc/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: usc.sh -i <file> [options]")
		os.Exit(1)
	}

	var configPath = flag.String("i", "", "Input configuration file")
	var outputPath = flag.String("o", "out", "Output directory for generated shell configs")
	var installFlag = flag.Bool("install", false, "Remove current configuration files and install generated shell configs")
	flag.Parse()

	configuration, err := config.LoadFromPath(context.Background(), *configPath)
	if err != nil {
		panic(err)
	}

	build.BuildTemplates(configuration, *outputPath, *installFlag)
}
