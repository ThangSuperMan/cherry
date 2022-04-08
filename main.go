package main

import (
	"fmt"
	"log"

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
	Commands: []*Z.Cmd{
		Bar,
	},
}

var Bar = &Z.Cmd{
	Name:    `bar`,
	Summary: `bar`,
	Call: func(_ *Z.Cmd, args ...string) error {
		fmt.Println("TESSTEEEEEE")
    return nil
	},
}
