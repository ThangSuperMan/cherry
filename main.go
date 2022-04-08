package main

import (
	"fmt"
	"log"

	"github.com/cherryramatisdev/cherry/git"
	"github.com/cherryramatisdev/cherry/install"
	"github.com/cherryramatisdev/cherry/monitor"
	"github.com/cherryramatisdev/cherry/tmux"
	"github.com/cherryramatisdev/cherry/twitch"
	Z "github.com/rwxrob/bonzai/z"
)

func main() {
	log.SetFlags(0)
	Cmd.Run()
}

var Cmd = &Z.Cmd{
	Name:      `cherry`,
	Summary:   `cherry's bonzai command tree`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2022 Cherry Ramatis`,
	License:   `Apache-2.0`,
	Call: func(_ *Z.Cmd, args ...string) error {
		fmt.Println(`
    Welcome to this personal assistant/collection of scripts

    Sub-Commands:
    - tmux
    - git(hub)
    - twitch
    - monitor
    - install
    `)
		return nil
	},
	Commands: []*Z.Cmd{
		tmux.Cmd,
		git.Cmd,
		twitch.Cmd,
		monitor.Cmd,
		install.Cmd,
	},
}
