package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// public variable
var HiCmd cli.Command

func init() {

	HiCmd = cli.Command{
		Name:  "finder",
		Usage: "find an package that's been registered",
		Action: func(ctx *cli.Context) error {
			fmt.Println("package found")
			return nil
		},
	}

}

func main() {
	app := &cli.App{}
	app.Commands = append(app.Commands, &HiCmd)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
