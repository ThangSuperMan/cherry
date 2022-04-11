package git

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/fs/file"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `git`,
	Summary:  `git command tree`,
	Aliases:  []string{"g"},
	Commands: []*Z.Cmd{help.Cmd, push},
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

		branch := currentBranch()

		fmt.Printf("git push origin %s", branch)
		return Z.Exec("git", "push", "origin", branch)
	},
}

func verifyJavascriptProject() {
	runShell("yarn", "build")
	runShell("yarn", "lint")
	runShell("yarn", "test")
	fmt.Print("\033[H\033[2J")
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

func currentBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Panic(err)
	}

	// Execute the command
	if err := cmd.Start(); err != nil {
		log.Panic(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(stdout)

	return strings.Replace(buf.String(), "\n", "", -1)
}
