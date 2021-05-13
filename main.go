package main

import (
	"fmt"

	"gitlab.com/altiano/goreen-tea/di"
	"gitlab.com/altiano/goreen-tea/src/frameworks/trace"
	"gitlab.com/altiano/goreen-tea/src/shared"
)

func main() {
	// config
	config := shared.NewConfig()

	// wire DI
	app := di.InjectApp(config)
	cli := di.InjectDummyCLI(config, app)
	rest := di.InjectIris(config, app)

	fmt.Println("Running... 🔥")

	go cli.Run()
	rest.Run()

	//
	cleanUp()
}

func cleanUp() {
	trace.Close()
}
