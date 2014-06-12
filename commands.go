package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandGen,
	commandList,
}

var commandGen = cli.Command{
	Name:        "gen",
	Usage:       "TODO",
	Description: "TODO",
	Action:      generate,
	Flags:       []cli.Flag{},
}

var commandList = cli.Command{
	Name:        "list",
	Usage:       "TODO",
	Description: "TODO",
	Action:      list,
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

func list(c *cli.Context) {
	args := c.Args()
	branchName := args.Get(0)
	path := args.Get(1)
	if path == "" {
		path = "./"
	}
	files, err := filepath.Glob(path + "/" + branchName + "~*")

	if err != nil {
		// TODO
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
