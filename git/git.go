package git

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
)

var Cmd = &Z.Cmd{
	Name:    `git`,
	Summary: `git command tree`,
	Call: func(_ *Z.Cmd, args ...string) error {
		fmt.Println(`git branch that will perform common scripting with tmux like`)
		return nil
	},
}
