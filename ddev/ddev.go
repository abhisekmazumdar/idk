package ddev

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Run a ddev command after replacing "idk" with "ddev"
func Run(cmd string) {
	cmd = replaceIDKWithDdev(cmd)
	command := exec.Command("ddev", cmd)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// Replace "idk" with "ddev"
func replaceIDKWithDdev(cmd string) string {
	return strings.Replace(cmd, "idk", "ddev", -1)
}
