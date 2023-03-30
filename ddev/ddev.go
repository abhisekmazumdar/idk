package ddev

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Run a ddev command after replacing "idk" with "ddev"
func Run(args string) {
	fmt.Println("Running: ddev " + args)
	command := exec.Command("ddev", strings.TrimSpace(args))
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Println(err)
	}
}
