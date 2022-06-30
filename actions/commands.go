package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Loud = cli.Command{
	Name:  "loud",
	Usage: "Send a loud message!",
	Action: func(c *cli.Context) error {
		fmt.Println("WARNING HEADPHONE USERS!! EXTREMELY LOUD CONTENT!!")
		return nil
	},
}

var Quiet = cli.Command{
	Name:  "quiet",
	Usage: "Returns a mouse's whisper!",
	Action: func(c *cli.Context) error {
		fmt.Println("shhhh... don't be too loud.")
		return nil
	},
}

var Special = cli.Command{
	Name:  "special",
	Usage: "A very special message for a very special person.",
	Action: func(c *cli.Context) error {
		fmt.Println("This is our most desperate hour. Help me, Obi-Wan Kenobi. You're my only hope.")
		return nil
	},
}
