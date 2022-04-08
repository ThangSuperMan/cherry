package monitor

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
)

var Cmd = &Z.Cmd{
	Name:    `monitor`,
	Summary: `monitor command tree`,
	Call: func(_ *Z.Cmd, args ...string) error {
		fmt.Println(`monitor/display management related commands`)
		return nil
	},
}
