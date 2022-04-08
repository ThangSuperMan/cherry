package twitch

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
)

var Cmd = &Z.Cmd{
	Name:    `twitch`,
	Summary: `twitch command tree`,
	Call: func(_ *Z.Cmd, args ...string) error {
		fmt.Println(`twitch related commands`)
		return nil
	},
}
