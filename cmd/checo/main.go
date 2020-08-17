package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var tag string
var ref string

func main() {
	os.Exit(run())
}

func run() int {
	app := cli.App{
		Copyright: "gophertribe",
		Name:      "checo",
		Authors: []*cli.Author{
			{Name: "Michal Klimuk", Email: "michal@gophertribe.com"},
		},
		Version: fmt.Sprintf("%s-%s", tag, ref),
		Commands: []*cli.Command{
			GameCmd, BoardCmd,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	return 0
}
