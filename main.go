package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var Version string = "HEAD"

func main() {
	app := cli.NewApp()
	app.Name = "migr"
	app.Usage = "Utility for SQL schema migration"
	app.Version = Version
	app.Author = "moznion"
	app.Email = "moznion@gmail.com"
	app.Commands = Commands
	app.Run(os.Args)
}
