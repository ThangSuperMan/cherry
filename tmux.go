package tmux

import "fmt"

var Cmd = &Z.Cmd{
	Name: `tmux`,
	Summary: `tmux branch that will perform common scriptings with tmux like:
  - Sending commands to chat on another tmux session
  - Setup the whole tmux setup with multiple sessions 
  - Kill all possible tmux session opened
  - Another tips and tricks
  `,
	Commands: []*Z.Cmd{
		foo,
	},
}

var foo = &Z.Cmd{
	Name:    `foo`,
	Summary: `foo`,
	Call: func(_ *Z.Cmd, args ...string) error {
		fmt.Println("foo")
		return nil
	},
}
