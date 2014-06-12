package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandGen,
}

var commandGen = cli.Command{
	Name:        "gen",
	Usage:       "TODO",
	Description: "TODO",
	Action:      generate,
	Flags:       []cli.Flag{},
}

func generate(c *cli.Context) {
	args := c.Args()
	sqlCommand := args.Get(0)
	tableName := args.Get(1)
	description := args.Get(2)

	if sqlCommand == "" || tableName == "" {
		cli.ShowCommandHelp(c, "gen")
		os.Exit(1)
	}

	branchName, err := getBranch()
	if err != nil {
		// TODO
		os.Exit(1)
	}

	datetime := time.Now().Format("20060102150405") // YYYYMMDDhhmmss
	filename := fmt.Sprintf("%s~%s-%s-%s", branchName, datetime, sqlCommand, tableName)
	if description != "" {
		filename += "-" + strings.Replace(description, " ", "_", -1)
	}
	filename += ".sql"

	sqlStatement := dumpSQLStatement(sqlCommand, tableName)

	err = ioutil.WriteFile(filename, []byte(sqlStatement), 0644)

	if err != nil {
		// TODO
		os.Exit(1)
	}
}
