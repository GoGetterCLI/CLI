package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// public variable
var HiCmd, Loud, Quiet, Special, Subcommands cli.Command

func init() {

	HiCmd = cli.Command{
		Name:  "finder",
		Usage: "find an package that's been registered",
		Action: func(ctx *cli.Context) error {
			fmt.Println("package found")
			return nil
		},
	}
	Loud = cli.Command{
		Name:  "loud",
		Usage: "Send a loud message!",
		Action: func(c *cli.Context) error {
			fmt.Println("WARNING HEADPHONE USERS!! EXTREMELY LOUD CONTENT!!")
			return nil
		},
	}

	Quiet = cli.Command{
		Name:  "quiet",
		Usage: "Returns a mouse's whisper!",
		Action: func(c *cli.Context) error {
			fmt.Println("shhhh... don't be too loud.")
			return nil
		},
	}

	Special = cli.Command{
		Name:  "special",
		Usage: "A very special message for a very special person.",
		Action: func(c *cli.Context) error {
			fmt.Println("This is our most desperate hour. Help me, Obi-Wan Kenobi. You're my only hope.")
			return nil
		},
	}

	Subcommands = cli.Command{
		Name:   "subcommand",
		Usage:  "Used to test subcommands using a specific library",
		Action: nil,
		Subcommands: []*cli.Command{
			{
				Name:  "pog",
				Usage: "Pog back at you!",
				Action: func(c *cli.Context) error {
					fmt.Println("foo bar test test 1234567890")
					return nil
				},
			},

			{
				Name:  "secondsubcommand",
				Usage: "Tests second subcommand",
				Action: func(c *cli.Context) error {
					fmt.Println("foo bar test test 1234567890")
					return nil
				},
			},
		},
	}

}

func main() {
	app := &cli.App{}
	app.Commands = append(app.Commands, &HiCmd, &Special, &Quiet, &Loud)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
