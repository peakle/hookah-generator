package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	// Version current tag version
	Version = "0"
	// CommitID of tag
	CommitID = "0"
	commands = []cli.Command{
		{
			Name:        "hg-server",
			ShortName:   "gstore",
			Description: "start server for parse apps",
			Action:      server.StartServer,
			Category:    "server",
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "hookah-generator"
	app.Commands = commands
	app.Version = fmt.Sprintf("%s - %s", Version, CommitID)

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error on run: " + err.Error())
	}
}
