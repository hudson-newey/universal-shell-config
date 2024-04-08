package main

import (
	"context"

	"usc/build"
	"usc/config"
)

func main() {
	configuration, err := config.LoadFromPath(context.Background(), "pkl/Config.pkl")
	if err != nil {
		panic(err)
	}

	build.BuildTemplate(configuration, "out")
}
