package lando

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Run a Lando command after replacing "idk" with "lando"
func Run(cmd string) {
	cmd = replaceIDKWithLando(cmd)
	command := exec.Command("lando", cmd)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// Replace "idk" with "lando"
func replaceIDKWithLando(cmd string) string {
	return strings.Replace(cmd, "idk", "lando", -1)
}
