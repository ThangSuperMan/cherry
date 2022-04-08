package install

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
)

var Cmd = &Z.Cmd{
	Name:    `install`,
	Summary: `install command tree`,
	Call: func(_ *Z.Cmd, args ...string) error {
		fmt.Println(`install packages related commands`)
		return nil
	},
}
