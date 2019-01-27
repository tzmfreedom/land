package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/urfave/cli.v1"
)

var Version string

func main() {
	godotenv.Load()

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(Version)
	}
	app := cli.NewApp()
	app.Name = "land"
	app.Usage = "Salesforce Apex Execution Environment on Local System"
	app.Version = Version
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
		visualforceCommand,
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
