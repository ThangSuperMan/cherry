package git

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/fs/file"
)

var Cmd = &Z.Cmd{
	Name:     `git`,
	Summary:  `git command tree`,
	Aliases:  []string{"g"},
	Commands: []*Z.Cmd{push},
}

var push = &Z.Cmd{
	Name:    `push`,
	Summary: `does git push`,
	Usage:   `push`,
	Aliases: []string{"p"},
	Call: func(_ *Z.Cmd, args ...string) error {
		if file.Exists("package.json") {
			fmt.Println("Loading...")
			verifyJavascriptProject()
		}

		fmt.Print("\033[H\033[2J")
		return Z.Exec("git", "commit", "--no-verify")
	},
}

func verifyJavascriptProject() {
	runShell("yarn", "build")
	runShell("yarn", "lint")
	runShell("yarn", "test")
}

func runShell(name string, args ...string) {
	cmd := exec.Command(name, args...)
	err, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(err)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
