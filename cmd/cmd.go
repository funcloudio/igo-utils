// Package cmd provides some functions for executing commands in Go programs.
package cmd

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func _exec(comm string, args ...string) {
	c := exec.Command(comm, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Exec executes a system command. It just wraps the standard package "os/exec".
func Exec(comm string) {
	if comm == "" {
		log.Println("A blank command")
		return
	}
	cs := strings.Split(comm, " ")
	_exec(cs[0], cs[1:]...)
}

// GoGet acts as "go get pkg_a pkg_b" command to install Go packages when you run GoGet(pkg_a, pkg_b).
// You also can call it as GoGet("-u", pkg_a) that's equal to "go get -u ...", other available
// arguments of "go get" can be passed in the same way.
func GoGet(pkgs ...string) {
	if len(pkgs) < 1 {
		log.Println("Not pass in any package name")
		return
	}
	_exec("go", append([]string{"get"}, pkgs...)...)
}
