package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	godotenv.Load()

	app := cli.NewApp()
	app.Commands = []cli.Command{
		dbSetupCommand,
		dbCreateCommand,
		dbSeedCommand,
		dbFetchCommand,
		testCommand,
		watchCommand,
		serverCommand,
		evalServerCommand,
		formatCommand,
		runCommand,
		checkCommand,
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
