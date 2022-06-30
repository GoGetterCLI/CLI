package actions

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
